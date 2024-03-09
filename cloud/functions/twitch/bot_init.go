package twitch

import (
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func BotInitEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	botInit(ctx, w, r)
}

func botInit(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	helixClient, err := providers.GetHelixAppClient(ctx)
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	if code == "" {
		url := helixClient.GetAuthorizationURL(&helix.AuthorizationURLParams{
			ResponseType: "code",
			State:        "bot",
			Scopes: []string{"user:bot", "user:read:chat", "user:write:chat", "moderator:manage:announcements", "moderator:manage:banned_users",
				"moderator:manage:chat_messages", "moderator:read:chat_settings", "moderator:manage:chat_settings",
				"moderator:read:chatters", "moderator:read:followers", "moderator:read:shield_mode", "moderator:manage:shield_mode",
				"user:edit", "user:manage:whispers"},
		})

		if _, err := w.Write([]byte(url)); err != nil {
			log.Printf("Failed to write response | %v", err)
		}
		return
	}

	if state != "bot" {
		log.Printf("Invalid state | %v", state)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	token, err := helixClient.RequestUserAccessToken(code)
	if err != nil {
		log.Printf("Failed to request user access token | %v", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}
	helixClient.SetUserAccessToken(token.Data.AccessToken)

	users, err := helixClient.GetUsers(&helix.UsersParams{})
	if err != nil {
		log.Printf("Failed to get user | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	if len(users.Data.Users) != 1 {
		log.Printf("Invalid user | %v", users.Data.Users)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	user := users.Data.Users[0]

	if user.ID != platform.TwitchUid {
		log.Printf("Invalid bot account | %s != %s", user.ID, platform.TwitchUid)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	pbToken := &pb.Token{
		AccessToken:  token.Data.AccessToken,
		RefreshToken: token.Data.RefreshToken,
	}

	data, err := proto.Marshal(pbToken)
	if err != nil {
		log.Printf("Failed to marshal token | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	secretManager, ok := ctx.GetSecretManager()
	if !ok {
		log.Printf("Secret manager not there")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if err := secretManager.SetSecret(ctx.Ctx, providers.SecretBotToken, data); err != nil {
		log.Printf("Failed to set secret | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Failed to write response | %v", err)
	}
}
