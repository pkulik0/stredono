package user

import (
	"errors"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"regexp"
)

const urlRegex = "^https://[a-zA-Z0-9-]+(.[a-zA-Z0-9-]+)+(/[a-zA-Z0-9-./?%&=]*)?$"
const storageUrlRegex = "^https://[a-z]*storage.googleapis.com/[a-zA-Z0-9-]+(.[a-zA-Z0-9-]+)+(/[a-zA-Z0-9-./?%&=]*)?$"

func validateUser(user *pb.User, uid string) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if user.Uid != uid {
		return errors.New("uid does not match")
	}

	matched, err := regexp.Match(`^[a-zA-Z0-9_]{4,32}$`, []byte(user.Username))
	if err != nil || !matched {
		return errors.New("invalid username")
	}

	matched, err = regexp.Match(`^[a-zA-Z0-9_]{4,64}$`, []byte(user.DisplayName))
	if err != nil || !matched {
		return errors.New("invalid display name")
	}

	if len(user.Description) > 1000 {
		return errors.New("description too long")
	}

	matched, err = regexp.Match(storageUrlRegex, []byte(user.PictureUrl))
	if len(user.PictureUrl) > 0 && (err != nil || !matched) {
		return errors.New("invalid picture url")
	}

	matched, err = regexp.Match(urlRegex, []byte(user.Url))
	if len(user.Url) > 0 && (err != nil || !matched) {
		return errors.New("invalid url")
	}

	return nil
}

func EditEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		Auth:  true,
		DocDb: true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	edit(ctx, w, r)
}

func edit(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	token, ok := ctx.GetAuthToken(r)
	if !ok {
		log.Errorf("received request without token")
		http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("failed to read request body | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	user := &pb.User{}
	if err := proto.Unmarshal(body, user); err != nil {
		log.Errorf("failed to unmarshal user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if err = validateUser(user, token.UserId()); err != nil {
		log.Errorf("invalid user | %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, ok := ctx.GetDocDb()
	if !ok {
		log.Errorf("failed to get doc db client")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	docs, err := db.Collection("users").Where("Username", "==", user.Username).Where("Uid", "!=", token.UserId()).Limit(1).Get(ctx.Ctx).All()
	if err != nil {
		log.Errorf("failed to get user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	if len(docs) > 0 {
		log.Errorf("username already taken")
		http.Error(w, "username already taken", http.StatusBadRequest)
		return
	}

	_, err = db.Collection("users").Doc(token.UserId()).Set(ctx.Ctx, user, modules.DbOpts{})
	if err != nil {
		log.Errorf("failed to set user | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
