package twitch

import (
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetDataEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		Auth:          true,
		DocDb:         true,
		SecretManager: true,
		Helix:         true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	getData(ctx, w, r)
}

func getData(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	twitchUser, err := getTwitchUser(ctx)
	if err != nil {
		log.Errorf("Failed to get twitch user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	db, ok := ctx.GetDocDb()
	if !ok {
		log.Errorf("Failed to get db client")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = db.Collection("twitch_users").Doc(twitchUser.Id).Set(ctx.Ctx, twitchUser, modules.DbOpts{})
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

func getTwitchUser(ctx *providers.Context) (*pb.TwitchUser, error) {
	helixClient, ok := ctx.GetHelix()
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
