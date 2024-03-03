package authentication

import (
	"akozadaev/go_reports/db/model"
	"context"
	"errors"
	"net/http"
)

const User = "accounts"

type Authentication interface {
	Authenticate(r *http.Request) (*model.Account, error)
}

func GetUser(ctx context.Context) (*model.Account, error) {
	currentUser, ok := ctx.Value(User).(*model.Account)
	if !ok {
		return nil, errors.New("context has no user")
	}

	return currentUser, nil
}
