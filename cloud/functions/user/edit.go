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

func validateAlerts(alerts []*pb.Alert) error {
	// TODO: implement
	return nil
}

const urlRegex = "^(http|https)://[a-zA-Z0-9-]+(.[a-zA-Z0-9-]+)+(/[a-zA-Z0-9-./?%&=]*)?$"

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

	if user.MinimumAmount < 0 {
		return errors.New("invalid minimum amount")
	}

	if len(user.Description) > 1000 {
		return errors.New("description too long")
	}

	if user.Currency == pb.Currency_UNKNOWN || user.Currency == pb.Currency_BITS {
		return errors.New("invalid currency")
	}

	matched, err = regexp.Match(urlRegex, []byte(user.PictureUrl)) // verify domain, allow only
	log.Printf("pic url: %s", user.PictureUrl)
	if len(user.PictureUrl) > 0 && (err != nil || !matched) {
		return errors.New("invalid picture url")
	}

	matched, err = regexp.Match(urlRegex, []byte(user.Url))
	if len(user.Url) > 0 && (err != nil || !matched) {
		return errors.New("invalid url")
	}

	return validateAlerts(user.Alerts)
}

func EditEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.CreateContext(r.Context(), &providers.Config{
		Auth:  true,
		DocDb: true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	r = r.WithContext(ctx)

	edit(w, r)
}

func edit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, _, ok := providers.GetAuthToken(ctx, r)
	if !ok {
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

	db, ok := providers.GetDocDb(ctx)
	if !ok {
		log.Errorf("failed to get doc db client")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = db.Collection("users").Doc(token.UserId()).Set(ctx, user, &modules.DbOpts{})
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
