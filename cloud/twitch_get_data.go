package cloud

import (
	"context"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func TwitchGetData(w http.ResponseWriter, r *http.Request) {
	platform.CorsMiddleware(platform.CloudMiddleware(&platform.CloudConfig{
		Auth: platform.AuthConfig{
			Client:      true,
			VerifyToken: true,
		},
		DocDb:          true,
		SecretsManager: true,
	}, platform.HelixMiddleware(twitchGetData)))(w, r)
}

func twitchGetData(w http.ResponseWriter, r *http.Request) {
	twitchUser, err := getTwitchUser(r.Context())
	if err != nil {
		log.Errorf("Failed to get twitch user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if err = saveTwitchUser(r.Context(), twitchUser); err != nil {
		log.Errorf("Failed to save twitch user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("Failed to write response | %s", err)
		return
	}
}

func getTwitchUser(ctx context.Context) (*pb.TwitchUser, error) {
	helixClient, ok := platform.GetHelix(ctx)
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	usersResponse, err := helixClient.GetUsers(&helix.UsersParams{})
	if err != nil {
		return nil, err
	}
	userData := usersResponse.Data.Users[0]

	return &pb.TwitchUser{
		Id:                userData.ID,
		DisplayName:       userData.DisplayName,
		Login:             userData.Login,
		AvatarUrl:         userData.ProfileImageURL,
		Description:       userData.Description,
		CreationTimestamp: userData.CreatedAt.Unix(),
	}, nil
}

func saveTwitchUser(ctx context.Context, user *pb.TwitchUser) error {
	db, ok := platform.GetDocDb(ctx)
	if !ok {
		return platform.ErrorMissingContextValue
	}

	_, err := db.Collection("twitch_users").Doc(user.Id).Set(ctx, user, &platform.DbOpts{})
	return err
}
