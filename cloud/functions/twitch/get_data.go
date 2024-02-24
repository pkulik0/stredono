package twitch

import (
	"context"
	"errors"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetDataEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.CreateContext(r.Context(), &providers.Config{
		Auth:          true,
		DocDb:         true,
		SecretManager: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	ctx, err = providers.CreateHelixContext(ctx, r)
	if err != nil {
		if errors.Is(err, platform.ErrorMissingContextValue) {
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	getData(w, r)
}

func getData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	twitchUser, err := getTwitchUser(ctx)
	if err != nil {
		log.Errorf("Failed to get twitch user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	db, ok := providers.GetDocDb(ctx)
	if !ok {
		log.Errorf("Failed to get db client")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = db.Collection("twitch_users").Doc(twitchUser.Id).Set(ctx, twitchUser, &modules.DbOpts{})
	if err != nil {
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
	helixClient, ok := providers.GetHelix(ctx)
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
