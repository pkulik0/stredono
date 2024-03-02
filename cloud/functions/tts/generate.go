package tts

import (
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const maxTextLength = 500

func GenerateEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		TextToSpeech:  true,
		Storage:       true,
		DocDb:         true,
		SecretManager: true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	generate(ctx, w, r)
}

func validateSpeechRequest(req *pb.SpeechRequest) error {
	if req.Id == "" {
		return platform.ErrorInvalidPayload
	}
	if req.Uid == "" {
		return platform.ErrorInvalidPayload
	}
	if req.VoiceIdBasic == "" {
		return platform.ErrorInvalidPayload
	}
	if req.VoiceIdPlus == "" {
		return platform.ErrorInvalidPayload
	}
	if req.Text == "" || len(req.Text) > maxTextLength {
		return platform.ErrorInvalidPayload
	}
	return nil
}

func handleGenerate(ctx *providers.Context, req *pb.SpeechRequest) (string, error) {
	storage, ok := ctx.GetStorage()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}
	bucket, err := storage.DefaultBucket()
	if err != nil {
		return "", err
	}

	ttsPlus, ok := ctx.GetTTSPlus()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}
	ttsBasic, ok := ctx.GetTTSBasic()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}

	audioData, err := ttsPlus.GenerateSpeech(ctx.Ctx, req.VoiceIdPlus, req.Text)
	if err != nil {
		audioData, err = ttsBasic.GenerateSpeech(ctx.Ctx, req.VoiceIdBasic, req.Text)
		if err != nil {
			return "", err
		}
	}

	path := fmt.Sprintf("users/%s/tts/%s-%d.mp3", req.Uid, req.Id, time.Now().Unix())
	obj := bucket.Object(path)
	wr := obj.NewWriter(ctx.Ctx)
	defer func(wr modules.Writer) {
		if err := wr.Close(); err != nil {
			log.Errorf("failed to close writer | %s", err)
		}
	}(wr)

	if _, err = wr.Write(audioData); err != nil {
		return "", err
	}
	return path, nil
}

func generate(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	//body, err := io.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	log.Errorf("failed to read request body | %s", err)
	//	http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
	//	return
	//}
	//req := pb.SpeechRequest{}
	//if err := proto.Unmarshal(body, &req); err != nil {
	//	log.Errorf("failed to unmarshal request | %s", err)
	//	http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
	//	return
	//}
	req := pb.SpeechRequest{
		Id:           "id",
		Uid:          "uid",
		VoiceIdPlus:  "5Q0t7uMcjvnagumLfvZi",
		VoiceIdBasic: "voiceIdBasic",
		Text:         "text",
	}

	if err := validateSpeechRequest(&req); err != nil {
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	path, err := handleGenerate(ctx, &req)
	if err != nil {
		log.Errorf("failed to handle generate | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(path)); err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
