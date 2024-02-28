package adapters

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
)

type GoogleTTS struct {
	Client *texttospeech.Client
}

func (tts *GoogleTTS) ProviderName() string {
	return "google"
}

func (tts *GoogleTTS) GenerateSpeech(ctx context.Context, voice string, text string) ([]byte, error) {
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			Name:         voice,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	speech, err := tts.Client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return nil, err
	}

	return speech.AudioContent, nil
}

func (tts *GoogleTTS) ListVoices(ctx context.Context, language string) ([]*pb.Voice, error) {
	response, err := tts.Client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{
		LanguageCode: language,
	})
	if err != nil {
		return nil, err
	}

	voices := make([]*pb.Voice, len(response.Voices))
	for i, voice := range response.Voices {
		voices[i] = &pb.Voice{
			Id:   voice.Name,
			Name: voice.Name,
		}
	}

	return voices, nil
}

func (tts *GoogleTTS) Close() error {
	return tts.Client.Close()
}
