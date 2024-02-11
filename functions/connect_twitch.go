package functions

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
	"net/http"
)

const (
	twitchClientId         = "t1kl0vkt6hv06bi4ah4691hi8fexso"
	twitchClientSecretName = "projects/621885503876/secrets/twitch-client-secret/versions/1"
	twitchRedirectUrl      = "http://localhost:8080/connectTwitchCallback"
)

func connectTwitch(w http.ResponseWriter, r *http.Request) {
	if setupCors(w, r) {
		return
	}

	Middleware(MiddlewareConfig{
		Auth: AuthConfig{
			Client: true,
			Token:  true,
		},
		Firestore: true,
		Secrets:   true,
	}, connectTwitchInternal)(w, r)
}

func connectTwitchCallback(w http.ResponseWriter, r *http.Request) {
	if setupCors(w, r) {
		return
	}

	Middleware(MiddlewareConfig{
		Firestore: true,
		Secrets:   true,
	}, connectTwitchCallbackInternal)(w, r)
}

func getTwitchOauthConfig(r *http.Request) (*oauth2.Config, error) {
	secretClient, ok := GetSecretsManager(r.Context())
	if !ok {
		return nil, errors.New("failed to get secrets manager")
	}

	clientSecretResponse, err := secretClient.AccessSecretVersion(r.Context(), &secretmanagerpb.AccessSecretVersionRequest{
		Name: twitchClientSecretName,
	})
	if err != nil {
		return nil, err
	}
	clientSecret := string(clientSecretResponse.Payload.Data)

	return &oauth2.Config{
		ClientID:     twitchClientId,
		ClientSecret: clientSecret,
		RedirectURL:  twitchRedirectUrl,
		Scopes: []string{
			"user:read:email",
			"channel:read:subscriptions",
		},
		Endpoint: twitch.Endpoint,
	}, nil
}

type TokenRequest struct {
	State       string `json:"state"`
	RedirectUrl string `json:"redirectUrl"`
}

func connectTwitchInternal(w http.ResponseWriter, r *http.Request) {
	token, ok := GetAuthToken(r.Context())
	if !ok {
		log.Error("Failed to get auth token")
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	redirectUrl := r.URL.Query().Get("redirect")
	if redirectUrl == "" {
		http.Error(w, "Missing redirect url", http.StatusBadRequest)
		return
	}

	twitchOauthConfig, err := getTwitchOauthConfig(r)
	if err != nil {
		log.Errorf("Failed to get twitch oauth config: %s", err)
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	firestoreClient, ok := GetFirestore(r.Context())
	if !ok {
		log.Error("Failed to get firestore client")
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	state := token.UID[:8] + "-" + uuid.New().String()[:8]
	req := TokenRequest{
		State:       state,
		RedirectUrl: redirectUrl,
	}
	_, err = firestoreClient.Collection("tokens").Doc("state").Set(r.Context(), map[string]interface{}{
		"twitch": map[string]interface{}{
			token.UID: req,
		},
	}, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to save state: %s", err)
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	url := twitchOauthConfig.AuthCodeURL(state)
	_, err = fmt.Fprintf(w, "%s", url)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}

func connectTwitchCallbackInternal(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}
	if state == "" {
		http.Error(w, "Missing state", http.StatusBadRequest)
		return
	}

	firestoreClient, ok := GetFirestore(r.Context())
	if !ok {
		log.Error("Failed to get firestore client")
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	statesDoc := firestoreClient.Collection("tokens").Doc("state")
	states, err := statesDoc.Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get states: %s", err)
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	twitchStates, ok := states.Data()["twitch"].(map[string]interface{})
	if !ok {
		log.Error("Failed to get twitch states")
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	uid := ""
	redirectUrl := ""
	log.Infof("States: %+v", twitchStates)
	for stateUid, s := range twitchStates {
		jsonData, err := json.Marshal(s)
		if err != nil {
			log.Errorf("Failed to marshal state: %s", err)
			http.Error(w, internalServerError, http.StatusInternalServerError)
			return
		}

		var ss TokenRequest
		err = json.Unmarshal(jsonData, &ss)
		if err != nil {
			log.Errorf("Failed to unmarshal state: %s", err)
			http.Error(w, internalServerError, http.StatusInternalServerError)
			return
		}

		log.Infof("State: %s, ss: %+v", state, ss)
		if state == ss.State {
			uid = stateUid
			redirectUrl = ss.RedirectUrl
			break
		}
	}
	if uid == "" {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	oauthConfig, err := getTwitchOauthConfig(r)
	if err != nil {
		log.Errorf("Failed to get twitch oauth config: %s", err)
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		log.Errorf("Failed to exchange code: %s", err)
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	_, err = firestoreClient.Collection("tokens").Doc(uid).Set(r.Context(), map[string]interface{}{
		"twitch": token,
	}, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to save token: %s", err)
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	_, err = statesDoc.Set(r.Context(), map[string]interface{}{
		"twitch": map[string]interface{}{
			uid: firestore.Delete,
		},
	}, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to delete state: %s", err)
	}

	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}
