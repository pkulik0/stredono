package tts

import (
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func UpdateEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		TextToSpeech: true,
		DocDb:        true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	update(ctx, w, r)
}

func updateSavedVoices(ctx *providers.Context) error {
	tts, ok := ctx.GetTTS()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	ttsVoices, err := tts.ListVoices(ctx.Ctx, "en-US")
	if err != nil {
		return err
	}

	db, ok := ctx.GetDocDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	provider := tts.ProviderName()
	log.Infof("updating update for provider %s", provider)
	_, err = db.Collection("tts").Doc(provider).Set(ctx.Ctx, &pb.TTSProvider{
		LastUpdated: time.Now().Unix(),
		Voices:      ttsVoices,
	}, modules.DbOpts{})
	if err != nil {
		return err
	}

	return nil
}

func update(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	if err := updateSavedVoices(ctx); err != nil {
		log.Errorf("failed to update saved update | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
