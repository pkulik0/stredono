package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"io"
	"net/http"
)

const basePath = "https://api.elevenlabs.io/v1/"

type ApiKeyTransport struct {
	Key  *string
	base http.RoundTripper
}

func (t *ApiKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("xi-api-key", *t.Key)
	req.Header.Set("Accept", "application/json")
	return t.base.RoundTrip(req)
}

type ElevenLabs struct {
	Client *http.Client
	apiKey string
}

func NewElevenLabs(apiKey string) *ElevenLabs {
	el := &ElevenLabs{
		Client: http.DefaultClient,
		apiKey: apiKey,
	}
	el.Client.Transport = &ApiKeyTransport{Key: &el.apiKey, base: http.DefaultTransport}
	return el
}

func (e *ElevenLabs) ProviderName() string {
	return "elevenlabs"
}

func (e *ElevenLabs) SetApiKey(apiKey string) {
	e.apiKey = apiKey
}

type elVoice struct {
	Id         string `json:"voice_id"`
	Name       string `json:"name"`
	PreviewUrl string `json:"preview_url"`
}

type elVoicesResponse struct {
	Voices []elVoice `json:"voices"`
}

func (e *ElevenLabs) ListVoices(ctx context.Context, language string) ([]*pb.Voice, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, basePath+"voices", nil)
	if err != nil {
		return nil, err
	}

	response, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var voicesResponse elVoicesResponse
	if err := json.NewDecoder(response.Body).Decode(&voicesResponse); err != nil {
		return nil, err
	}

	voices := make([]*pb.Voice, len(voicesResponse.Voices))
	for i, voice := range voicesResponse.Voices {
		voices[i] = &pb.Voice{
			Id:        voice.Id,
			Name:      voice.Name,
			SampleUrl: voice.PreviewUrl,
		}
	}

	return voices, nil
}

func (e *ElevenLabs) GenerateSpeech(ctx context.Context, voice string, text string) ([]byte, error) {
	data := map[string]string{
		"model_id": "eleven_multilingual_v2",
		"text":     text,
	}
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(data); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%stext-to-speech/%s?output_format=mp3_44100_128", basePath, voice)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	response, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}
