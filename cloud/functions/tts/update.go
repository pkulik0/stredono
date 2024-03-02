package tts

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func UpdateEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		TextToSpeech:  true,
		DocDb:         true,
		SecretManager: true,
		Storage:       true,
		Proxy:         true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	update(ctx, w, r)
}

func updateSavedVoices(ctx *providers.Context) error {
	ttsBasic, ok := ctx.GetTTSBasic()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	ttsPlus, ok := ctx.GetTTSPlus()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	db, ok := ctx.GetDocDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	storage, ok := ctx.GetStorage()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	namesSnap, err := db.Collection("tts").Doc("voice-names").Get(ctx.Ctx)
	if err != nil {
		return err
	}
	var names map[string]string
	if err := namesSnap.DataTo(&names); err != nil {
		return err
	}

	updateVoices := func(tts modules.TTS) ([]*pb.Voice, error) {
		ttsVoices, err := tts.GetVoices(ctx.Ctx)
		if err != nil {
			return nil, err
		}

		namedVoices := make([]*pb.Voice, 0)
		for _, voice := range ttsVoices {
			name, ok := names[voice.Id]
			if !ok {
				continue
			}
			voice.Name = name
			namedVoices = append(namedVoices, voice)
		}

		return namedVoices, nil
	}

	basicVoices, err := updateVoices(ttsBasic)
	if err != nil {
		return err
	}
	plusVoices, err := updateVoices(ttsPlus)
	if err != nil {
		return err
	}

	var voices []*pb.Voice
	voices = append(voices, plusVoices...)
	voices = append(voices, basicVoices...)

	bucket, err := storage.DefaultBucket()
	if err != nil {
		return err
	}

	samplesSnap, err := db.Collection("tts").Doc("samples").Get(ctx.Ctx)
	if err != nil {
		return err
	}
	var samples map[string]string
	if err = samplesSnap.DataTo(&samples); err != nil {
		return err
	}

	for _, voice := range voices {
		var tts modules.TTS
		var sampleKey string

		switch voice.Tier {
		case pb.Tier_BASIC:
			tts = ttsBasic
			sampleKey = voice.Languages[0] // Basic voices have only one language
		case pb.Tier_PLUS:
			tts = ttsPlus
			sampleKey = "multi" // Reserved language code for plus (multilanguage) voices
		default:
			return platform.ErrorUnknownEnumValue
		}

		text, ok := samples[sampleKey]
		if !ok {
			log.Errorf("missing tts sample text for key %s", sampleKey)
			return platform.ErrorMissingSample
		}

		obj := bucket.Object(fmt.Sprintf("tts/samples/%s_%x.mp3", voice.Name, sha256.Sum256([]byte(text))))
		attrs, err := obj.Attrs(ctx.Ctx)
		if err == nil {
			voice.SampleUrl = attrs.MediaUrl
			continue
		}
		if !errors.Is(err, platform.ErrorObjectNotFound) {
			return err
		}

		sample, err := tts.GenerateSpeech(ctx.Ctx, voice.Id, text)
		if err != nil {
			return err
		}

		wr := obj.NewWriter(ctx.Ctx)
		if _, err = wr.Write(sample); err != nil {
			return err
		}
		if err = wr.Close(); err != nil {
			return err
		}

		if err := obj.SetPublicRead(ctx.Ctx); err != nil {
			return err
		}

		attrs, err = obj.Attrs(ctx.Ctx)
		if err != nil {
			return err
		}
		voice.SampleUrl = attrs.MediaUrl

		log.Infof("saved new sample for voice %s (%s)", voice.Id, voice.Name)
	}

	log.Infof("Found %d voices (basic: %d, plus: %d)", len(voices), len(basicVoices), len(plusVoices))
	_, err = db.Collection("tts").Doc("voices").Set(ctx.Ctx, &pb.Voices{Voices: voices}, modules.DbOpts{})
	return err
}

func update(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	if err := updateSavedVoices(ctx); err != nil {
		log.Errorf("failed to update tts voices | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
