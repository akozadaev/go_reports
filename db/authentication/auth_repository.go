package authentication

import (
	database "akozadaev/go_reports/db"
	"akozadaev/go_reports/db/model"
	"context"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByLogin(ctx context.Context, login string) model.Account
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) FindByLogin(ctx context.Context, login string) model.Account {
	db := database.FromContext(ctx, a.db)

	var user model.Account
	db.First(&user, "email = ?", login)

	return user
}
