package adapters

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"strings"
)

type GoogleTTS struct {
	Client *texttospeech.Client
}

func (tts *GoogleTTS) ProviderName() string {
	return "google"
}

func (tts *GoogleTTS) GenerateSpeech(ctx context.Context, voice string, text string) ([]byte, error) {
	voiceNameParts := strings.Split(voice, "-")
	if len(voiceNameParts) != 4 {
		return nil, platform.ErrorInvalidPayload
	}
	languageCode := strings.Join(voiceNameParts[:2], "-")

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: languageCode,
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

func (tts *GoogleTTS) GetVoices(ctx context.Context) ([]*pb.Voice, error) {
	response, err := tts.Client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{})
	if err != nil {
		return nil, err
	}

	allowedTypes := []string{"Wavenet", "Neural2", "Standard"}
	voices := make([]*pb.Voice, 0)
	for _, voice := range response.Voices {
		for _, allowedType := range allowedTypes {
			if !strings.Contains(voice.Name, allowedType) {
				continue
			}
		}

		languageCodes := make([]string, len(voice.LanguageCodes))
		for i, code := range voice.LanguageCodes {
			languageCodes[i] = strings.Split(code, "-")[0]
		}

		voices = append(voices, &pb.Voice{
			Id:        voice.Name,
			Name:      voice.Name,
			Languages: languageCodes,
			Tier:      pb.Tier_BASIC,
			Provider:  tts.ProviderName(),
			Gender:    genderToEnum(voice.SsmlGender.String()),
		})
	}

	return voices, nil
}

func (tts *GoogleTTS) Close() error {
	return tts.Client.Close()
}
