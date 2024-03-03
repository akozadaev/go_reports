package authentication

import (
	"akozadaev/go_reports/db/model"
	"akozadaev/go_reports/utils"
	"encoding/base64"
	"errors"

	/*	authRepository "go-aws-service/internal/app/database"
		"go-aws-service/internal/team/model"
		"go-aws-service/pkg/logging"
		"go-aws-service/pkg/utils"*/
	"net/http"
	"strings"
)

type basicAuth struct {
	repo AuthRepository
}

func NewBasicAuth(repo AuthRepository) Authentication {
	return &basicAuth{
		repo: repo,
	}
}

func (b *basicAuth) Authenticate(r *http.Request) (*model.Account, error) {
	if !b.HasAuthHeader(r) {
		return nil, errors.New("header_not_found")
	}

	header := r.Header.Get("Authorization")
	token, err := b.createTokenFromHeader(header)

	if err != nil {
		return nil, err
	}
	requestedUser, err := b.parseToken(token)
	if err != nil {
		return nil, err
	}

	user := b.repo.FindByLogin(r.Context(), requestedUser.Username)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPasswordHash(requestedUser.Password, user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (b *basicAuth) HasAuthHeader(r *http.Request) bool {
	header := r.Header.Get("Authorization")
	if header == "" {
		return false
	}

	parts := strings.Fields(header)
	if len(parts) != 2 || parts[0] != "Basic" {
		return false
	}

	return true
}

func (b *basicAuth) createTokenFromHeader(header string) (string, error) {
	parts := strings.Fields(header)
	if len(parts) != 2 || parts[0] != "Basic" {
		return "", errors.New("invalid_token")
	}

	token := parts[1]

	return token, nil
}

func (b *basicAuth) parseToken(token string) (*model.AuthUserRequest, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return nil, errors.New("invalid_token")
	}

	return &model.AuthUserRequest{
		Username: parts[0],
		Password: parts[1],
	}, nil
}
