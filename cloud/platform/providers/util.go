package providers

import (
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	log "github.com/sirupsen/logrus"
)

func ProviderIdToUid(ctx *Context, provider string, ID string) (string, error) {
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}

	uid := ""
	if err := rtdb.NewRef("Users").Child(provider).Child(ID).Child("Uid").Get(ctx.Ctx, &uid); err != nil {
		return "", err
	}

	if uid == "" {
		return "", fmt.Errorf("user not found")
	}

	return uid, nil
}

func GenerateSpeech(ctx *Context, req *pb.TTSRequest) (string, error) {
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

	var audioData []byte
	if req.Settings.Tier == pb.Tier_PLUS {
		audioData, err = ttsPlus.GenerateSpeech(ctx.Ctx, req.Settings.VoiceIdPlus, req.Text)
		if err != nil {
			log.Printf("Failed to generate speech with TTS+ | %v", err)
		}
	}

	// Means Tier != Plus or failed to generate speech with TTS+
	if audioData == nil {
		audioData, err = ttsBasic.GenerateSpeech(ctx.Ctx, req.Settings.VoiceIdBasic, req.Text)
		if err != nil {
			return "", err
		}
	}

	path := fmt.Sprintf("users/%s/tts/%s.mp3", req.Uid, req.ID)
	obj := bucket.Object(path)
	wr := obj.NewWriter(ctx.Ctx)
	if _, err = wr.Write(audioData); err != nil {
		return "", err
	}
	if err := wr.Close(); err != nil {
		return "", err
	}

	if err := obj.SetPublicRead(ctx.Ctx); err != nil {
		return "", err
	}

	attrs, err := obj.Attrs(ctx.Ctx)
	if err != nil {
		return "", err
	}

	return attrs.MediaUrl, nil
}
