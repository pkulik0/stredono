package providers

import (
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	log "github.com/sirupsen/logrus"
	"time"
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

	audioData, err := ttsPlus.GenerateSpeech(ctx.Ctx, req.Settings.VoiceIdPlus, req.Text)
	if err != nil {
		audioData, err = ttsBasic.GenerateSpeech(ctx.Ctx, req.Settings.VoiceIdBasic, req.Text)
		if err != nil {
			return "", err
		}
	}

	path := fmt.Sprintf("users/%s/tts/%s-%d.mp3", req.Uid, req.ID, time.Now().Unix())
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
