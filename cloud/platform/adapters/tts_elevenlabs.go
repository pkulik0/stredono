package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	basePath        = "https://api.elevenlabs.io/v1/"
	elevenlabsModel = "eleven_multilingual_v2"
)

var elevenlabsV2Languages = []string{
	"en", "ja", "zh", "de", "hi", "fr", "ko", "pt", "it", "es", "id", "nl", "tr", "fil", "pl", "sv", "bg", "ro", "ar",
	"cs", "el", "fi", "hr", "ms", "sk", "da", "ta", "uk", "ru",
}

type apiKeyTransport struct {
	Key  *string
	base http.RoundTripper
}

func (t *apiKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("xi-api-key", *t.Key)
	req.Header.Set("Accept", "application/json")
	return t.base.RoundTrip(req)
}

type ElevenLabs struct {
	client     *http.Client
	docDb      modules.DocDb
	collection string
	apiKey     string
	maxRetries int
}

func NewElevenLabs(docDb modules.DocDb, collection string) (*ElevenLabs, error) {
	el := &ElevenLabs{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
		docDb:      docDb,
		collection: collection,
		maxRetries: 3,
	}
	el.client.Transport = &apiKeyTransport{Key: &el.apiKey, base: http.DefaultTransport}
	return el, nil
}

func (e *ElevenLabs) ProviderName() string {
	return "elevenlabs"
}

func (e *ElevenLabs) SetApiKey(apiKey string) {
	e.apiKey = apiKey
}

type elLabels struct {
	Gender      string `json:"gender"`
	Age         string `json:"age"`
	Accent      string `json:"accent"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

type elVoice struct {
	Id         string `json:"voice_id"`
	Name       string `json:"name"`
	PreviewUrl string `json:"preview_url"`
	Labels     elLabels
}

type elVoicesResponse struct {
	Voices []elVoice `json:"voices"`
}

func genderToEnum(gender string) pb.Gender {
	switch strings.ToLower(gender) {
	case "male":
		return pb.Gender_MALE
	case "female":
		return pb.Gender_FEMALE
	default:
		return pb.Gender_NOT_SPECIFIED
	}
}

func (e *ElevenLabs) GetVoices(ctx context.Context) ([]*pb.Voice, error) {
	voices := make([]*pb.Voice, 0)
	err := e.getKeyAndRunWithRetry(ctx, 0, func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, basePath+"voices", nil)
		if err != nil {
			return err
		}

		response, err := e.client.Do(req)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		var voicesResponse elVoicesResponse
		if err := json.NewDecoder(response.Body).Decode(&voicesResponse); err != nil {
			return err
		}

		for _, voice := range voicesResponse.Voices {
			voices = append(voices, &pb.Voice{
				Id:        voice.Id,
				Name:      voice.Name,
				SampleUrl: voice.PreviewUrl,
				Languages: elevenlabsV2Languages,
				Tier:      pb.Tier_PLUS,
				Provider:  e.ProviderName(),
				Gender:    genderToEnum(voice.Labels.Gender),
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return voices, nil
}

func (e *ElevenLabs) getKeyAndRun(ctx context.Context, requiredCharacters int, f func(context.Context) error) error {
	return e.docDb.RunTransaction(ctx, func(ctx context.Context, tx modules.Transaction) error {
		iter := tx.Documents(e.docDb.Collection(e.collection).Where("CharactersLeft", ">",
			requiredCharacters).Limit(1))
		defer iter.Stop()
		snap, err := iter.Next()
		if err != nil {
			return err
		}

		var ttsKey *pb.TTSKey
		if err = snap.DataTo(&ttsKey); err != nil {
			return err
		}
		e.SetApiKey(ttsKey.Key)
		log.Printf("Using EL key with id %s", snap.Ref().Id())

		// Prevent skipping the update of the key data in db
		delayedErr := f(ctx)

		info, err := e.getUserInfo(ctx)
		if err != nil {
			return err
		}

		ttsKey.CharactersLeft = int32(info.CharacterLimit - info.CharacterCount)
		ttsKey.ResetTimestamp = info.NextResetTime
		ttsKey.LastUsed = time.Now().Unix()

		if err = tx.Set(snap.Ref(), ttsKey, modules.DbOpts{}); err != nil {
			if delayedErr != nil {
				return fmt.Errorf("%s | %s", delayedErr, err)
			}
			return err
		}

		return delayedErr
	})
}

func (e *ElevenLabs) getKeyAndRunWithRetry(ctx context.Context, requiredCharacters int, f func(context.Context) error) error {
	var err error
	for i := 0; i < e.maxRetries; i++ {
		err := e.getKeyAndRun(ctx, requiredCharacters, f)
		if err == nil {
			return nil
		}
		log.Errorf("failed to get key and run function, retrying with new key, attempt %d/%d\n%v", i+1, e.maxRetries, err)
	}
	return err
}

func (e *ElevenLabs) GenerateSpeech(ctx context.Context, voice string, text string) ([]byte, error) {
	var audio []byte
	err := e.getKeyAndRunWithRetry(ctx, len(text), func(ctx context.Context) error {
		data := map[string]string{
			"model_id": elevenlabsModel,
			"text":     text,
		}
		body := new(bytes.Buffer)
		if err := json.NewEncoder(body).Encode(data); err != nil {
			return err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%stext-to-speech/%s?output_format=mp3_44100_128", basePath, voice), body)
		if err != nil {
			return err
		}
		response, err := e.client.Do(req)
		if err != nil {
			return err
		}
		if response.StatusCode != http.StatusOK {
			return platform.ErrorResponseNotOk
		}

		defer response.Body.Close()
		audio, err = io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return audio, nil
}

type elUserInfo struct {
	CharacterCount int   `json:"character_count"`
	CharacterLimit int   `json:"character_limit"`
	NextResetTime  int64 `json:"next_character_count_reset_unix"`
}

func (e *ElevenLabs) getUserInfo(ctx context.Context) (*elUserInfo, error) {
	var userInfo *elUserInfo
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, basePath+"user/subscription", nil)
	if err != nil {
		return nil, err
	}

	response, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, platform.ErrorResponseNotOk
	}

	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
