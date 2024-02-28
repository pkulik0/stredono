package tts

import (
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
)

const maxTextLength = 500

func GenerateEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		TextToSpeech: true,
		Storage:      true,
	})
	if err != nil {
		log.Error(err)
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
	if req.VoiceId == "" {
		return platform.ErrorInvalidPayload
	}
	if req.Text == "" || len(req.Text) > maxTextLength {
		return platform.ErrorInvalidPayload
	}
	return nil
}

func generate(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	//defer r.Body.Close()
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Errorf("failed to read request body | %s", err)
	//	http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
	//	return
	//}
	//
	//req := pb.SpeechRequest{}
	//if err := proto.Unmarshal(body, &req); err != nil {
	//	log.Errorf("failed to unmarshal request | %s", err)
	//	http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
	//	return
	//}
	//
	//if err := validateSpeechRequest(&req); err != nil {
	//	http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
	//	return
	//}

	storage, ok := ctx.GetStorage()
	if !ok {
		log.Error("storage client not found")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	tts, ok := ctx.GetTTS()
	if !ok {
		log.Error("tts client not found")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	bucket, err := storage.DefaultBucket()
	if err != nil {
		log.Errorf("failed to get default bucket | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	voices, err := tts.ListVoices(ctx.Ctx, "en-US")
	if err != nil {
		log.Errorf("failed to list update | %s", err)
		return
	}

	audioData, err := tts.GenerateSpeech(ctx.Ctx, voices[rand.Intn(len(voices))].Id,
		"Adam wysłał 154.50 zł")
	if err != nil {
		log.Errorf("failed to generate speech | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	//path := "users/" + req.Uid + "/tts/" + req.Id + ".mp3"
	obj := bucket.Object("aha.mp3")
	wr := obj.NewWriter(ctx.Ctx)
	defer func(wr modules.Writer) {
		if err := wr.Close(); err != nil {
			log.Errorf("failed to close writer | %s", err)
		}
	}(wr)

	if _, err := wr.Write(audioData); err != nil {
		log.Errorf("failed to write audio data | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte("OK")); err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
