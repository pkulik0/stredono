package functions

import (
	"cloud.google.com/go/firestore"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	CorsMiddleware(CloudMiddleware(CloudConfig{
		Auth: AuthConfig{
			Client: true,
			Token:  true,
		},
		Firestore: true,
		Secrets:   true,
	}, HelixMiddleware(getData)))(w, r)
}

func getData(w http.ResponseWriter, r *http.Request) {
	helixClient, ok := GetHelixClient(r.Context())
	if !ok {
		log.Errorf("Failed to get twitch client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	token, ok := GetAuthToken(r.Context())
	if !ok {
		log.Error("Failed to get auth token")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	firestoreClient, ok := GetFirestore(r.Context())
	if !ok {
		log.Error("Failed to get firestore client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	usersResponse, err := helixClient.GetUsers(&helix.UsersParams{})
	if err != nil {
		log.Errorf("Failed to get user: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	userData := usersResponse.Data.Users[0]
	user := &pb.TwitchUser{
		Id:                userData.ID,
		DisplayName:       userData.DisplayName,
		Login:             userData.Login,
		AvatarUrl:         userData.ProfileImageURL,
		Description:       userData.Description,
		CreationTimestamp: userData.CreatedAt.Unix(),
	}

	rewardsResponse, err := helixClient.GetCustomRewards(&helix.GetCustomRewardsParams{
		BroadcasterID: user.Id,
	})
	if err != nil {
		log.Errorf("Failed to get rewards: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	rewardsData := rewardsResponse.Data.ChannelCustomRewards
	rewards := make([]*pb.TwitchReward, len(rewardsData))
	for i, rewardData := range rewardsData {
		rewards[i] = &pb.TwitchReward{
			Id:        rewardData.ID,
			Name:      rewardData.Prompt,
			Cost:      int64(rewardData.Cost),
			IsEnabled: rewardData.IsEnabled,
		}
	}

	_, err = firestoreClient.Collection("twitch").Doc(token.UID).Set(r.Context(), map[string]interface{}{
		"user":    user,
		"rewards": rewards,
	}, firestore.MergeAll)
	if err != nil {
		log.Errorf("Failed to save token: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Success"))
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
