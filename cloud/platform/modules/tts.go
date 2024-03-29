package modules

import (
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
)

type Voice struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Language string `json:"-"`
}

type TTS interface {
	GenerateSpeech(ctx context.Context, voice string, text string) ([]byte, error)
	GetVoices(ctx context.Context) ([]*pb.Voice, error)
	ProviderName() string
}
