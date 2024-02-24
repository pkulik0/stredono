package user

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/golang-jwt/jwt"
	"io"
	"net/http"
	"testing"
)

func newTestSignerKeys() (map[string]string, *rsa.PrivateKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	privPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes})

	return map[string]string{"test": string(privPem)}, priv, nil
}

func TestRequestToClaims(t *testing.T) {
	signingKeys, privKey, err := newTestSignerKeys()
	if err != nil {
		t.Fatalf("failed to create signing keys: %v", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "test",
	})
	tokenString, err := jwtToken.SignedString(privKey)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	reqBody := onRegisterRequest{Data: onRegisterPayload{
		JwtToken: tokenString,
	}}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := &http.Request{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(reqBodyBytes)),
	}

	_, err = requestToClaims(req, signingKeys)
	if err != nil {
		t.Fatalf("failed to get claims: %v", err)
	}
}
