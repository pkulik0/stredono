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
	twitchClientSecretName = GcSecretsPath + "/twitch-client-secret/versions/latest"
	twitchRedirectUrl      = "http://localhost:8080/connectTwitchCallback"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	CorsMiddleware(CloudMiddleware(CloudConfig{
		Auth: AuthConfig{
			Client: true,
			Token:  true,
		},
		Firestore: true,
		Secrets:   true,
	}, connect))(w, r)
}

func ConnectCallback(w http.ResponseWriter, r *http.Request) {
	CorsMiddleware(CloudMiddleware(CloudConfig{
		Firestore: true,
		Secrets:   true,
	}, connectCallback))(w, r)
}

func getOauthConfig(r *http.Request) (*oauth2.Config, error) {
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
			"moderator:read:followers", // channel.follow
			"channel:read:subscriptions",
			"channel:read:redemptions",
			"bits:read",
			"channel:manage:ads",
			"channel:read:ads", // channel.ad_break_begin
			"channel:manage:broadcast",
			"channel:edit:commercial",
			"channel:read:hype_train",
			"channel:read:goals",
			"channel:read:vips",
			"user:read:broadcast",
			"user:read:chat",
		},
		Endpoint: twitch.Endpoint,
	}, nil
}

type TokenRequest struct {
	State       string `json:"state"`
	RedirectUrl string `json:"redirectUrl"`
}

func connect(w http.ResponseWriter, r *http.Request) {
	token, ok := GetAuthToken(r.Context())
	if !ok {
		log.Error("Failed to get auth token")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	redirectUrl := r.URL.Query().Get("redirect")
	if redirectUrl == "" {
		http.Error(w, "Missing redirect url", http.StatusBadRequest)
		return
	}

	twitchOauthConfig, err := getOauthConfig(r)
	if err != nil {
		log.Errorf("Failed to get twitch oauth config: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	firestoreClient, ok := GetFirestore(r.Context())
	if !ok {
		log.Error("Failed to get firestore client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
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
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	url := twitchOauthConfig.AuthCodeURL(state)
	_, err = fmt.Fprintf(w, "%s", url)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}

func connectCallback(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	statesDoc := firestoreClient.Collection("tokens").Doc("state")
	states, err := statesDoc.Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get states: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	twitchStates, ok := states.Data()["twitch"].(map[string]interface{})
	if !ok {
		log.Error("Failed to get twitch states")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	uid := ""
	redirectUrl := ""
	log.Infof("States: %+v", twitchStates)
	for stateUid, s := range twitchStates {
		jsonData, err := json.Marshal(s)
		if err != nil {
			log.Errorf("Failed to marshal state: %s", err)
			http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		var ss TokenRequest
		err = json.Unmarshal(jsonData, &ss)
		if err != nil {
			log.Errorf("Failed to unmarshal state: %s", err)
			http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
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

	oauthConfig, err := getOauthConfig(r)
	if err != nil {
		log.Errorf("Failed to get twitch oauth config: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		log.Errorf("Failed to exchange code: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = firestoreClient.Collection("tokens").Doc(uid).Set(r.Context(), map[string]interface{}{
		"twitch": token,
	}, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to save token: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
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
