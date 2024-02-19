package functions

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/pkulik0/stredono/functions/pb"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"regexp"
)

func OnRegister(w http.ResponseWriter, r *http.Request) {
	CloudMiddleware(CloudConfig{
		Firestore: true,
		Auth: AuthConfig{
			Client: true,
		},
	}, onRegister)(w, r)
}

func onRegister(w http.ResponseWriter, r *http.Request) {
	googleKeys, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		log.Errorf("Failed to get google keys: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	var keys map[string]string
	if err = json.NewDecoder(googleKeys.Body).Decode(&keys); err != nil {
		log.Errorf("Failed to decode google keys: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed to read request: %s", err)
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Infof("OnRegister Body: %s", body)

	recvData := struct {
		Data struct {
			JwtToken string `json:"jwt"`
		} `json:"data"`
	}{}
	if err := json.Unmarshal(body, &recvData); err != nil {
		log.Errorf("Failed to unmarshal request: %s", err)
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	var token *jwt.Token
	for _, key := range keys {
		token, err = jwt.Parse(recvData.Data.JwtToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err == nil {
			break
		}
	}
	if token == nil {
		http.Error(w, "No valid token", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Failed to get claims", http.StatusBadRequest)
		return
	}

	log.Infof("OnRegister Claims: %s", claims)

	aud, ok := claims["aud"].(string)
	if !ok {
		log.Errorf("Failed to get audience")
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	matched, err := regexp.Match(`^https://onregister-[a-z0-9-.]+\.run\.app$`, []byte(aud))
	if err != nil {
		log.Errorf("Failed to match audience: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	if !matched {
		http.Error(w, "Invalid audience", http.StatusBadRequest)
		return
	}

	uid, ok := claims["sub"].(string)
	if !ok {
		log.Errorf("Failed to get subject")
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	user := &pb.User{
		Username:      "",
		Uid:           uid,
		Url:           "",
		Description:   "",
		MinimumAmount: 1,
		Currency:      pb.Currency_PLN,
		Alerts:        make([]*pb.Alert, 0),
	}

	firestoreClient, ok := GetFirestore(r.Context())
	if !ok {
		log.Errorf("Failed to get firestore client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	_, err = firestoreClient.Collection("users").Doc(uid).Create(r.Context(), user)
	if err != nil {
		log.Errorf("Failed to set user: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
