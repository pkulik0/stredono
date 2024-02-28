package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"io"
	"net/http"
	"time"
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
	client     *http.Client
	docDb      modules.DocDb
	collection string
	apiKey     string
}

func NewElevenLabs(docDb modules.DocDb, collection string) *ElevenLabs {
	el := &ElevenLabs{
		client:     http.DefaultClient,
		docDb:      docDb,
		collection: collection,
	}
	el.client.Transport = &ApiKeyTransport{Key: &el.apiKey, base: http.DefaultTransport}
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

	response, err := e.client.Do(req)
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
	q := e.docDb.Collection(e.collection).Where("CharactersLeft", ">", len(text))
	var audio []byte
	err := e.docDb.RunTransaction(ctx, func(ctx context.Context, tx modules.Transaction) error {
		iter := tx.Documents(q)
		defer iter.Stop()
		snap, err := iter.Next()
		if err != nil {
			return err
		}

		var ttsKey *pb.TTSKey
		if err := snap.DataTo(&ttsKey); err != nil {
			return err
		}
		e.SetApiKey(ttsKey.Key)

		data := map[string]string{
			"model_id": "eleven_multilingual_v2",
			"text":     text,
		}
		body := new(bytes.Buffer)
		if err := json.NewEncoder(body).Encode(data); err != nil {
			return err
		}

		url := fmt.Sprintf("%stext-to-speech/%s?output_format=mp3_44100_128", basePath, voice)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
		if err != nil {
			return err
		}

		response, err := e.client.Do(req)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		audio, err = io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		info, err := e.GetUserInfo(ctx)
		if err != nil {
			return err
		}

		ttsKey.CharactersLeft = ttsKey.CharactersLeft - int32(len(text))
		ttsKey.ResetTimestamp = info.NextResetTime
		ttsKey.LastUsed = time.Now().Unix()

		if err = tx.Set(snap.Ref(), ttsKey, modules.DbOpts{}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return audio, nil
}

type UserInfo struct {
	CharacterCount int   `json:"character_count"`
	CharacterLimit int   `json:"character_limit"`
	NextResetTime  int64 `json:"next_character_count_reset_unix"`
}

func (e *ElevenLabs) GetUserInfo(ctx context.Context) (*UserInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, basePath+"user/subscription", nil)
	if err != nil {
		return nil, err
	}

	response, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func (e *ElevenLabs) SaveUserInfo(ctx context.Context, userInfo *UserInfo) error {
	_, err := e.docDb.Collection(e.collection).Doc("elevenlabs").Set(ctx, map[string]interface{}{
		"status": map[string]interface{}{
			e.apiKey: userInfo,
		},
	}, modules.DbOpts{MergeAll: true})

	return err
}
