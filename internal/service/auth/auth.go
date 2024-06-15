package auth

import (
	"SingleSignOnService/internal/domain/model"
	"context"
	"log/slog"
	"time"
)

type Auth struct {
	log       *slog.Logger
	uProvider UserProvider
	uSaver    UserSaver
	aProvider AppProvider
	tokenTTL  time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		hashedPassword string,
	) (uuid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (model.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int64) (model.App, error)
}

func New(
	log *slog.Logger,
	userProvider UserProvider,
	saver UserSaver,
	appProvider AppProvider,
	tokenTTL time.Duration) *Auth {
	return &Auth{
		log:       log,
		uProvider: userProvider,
		uSaver:    saver,
		aProvider: appProvider,
		tokenTTL:  tokenTTL,
	}
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, error) {
	panic("not implemented")
}

func (a *Auth) SignUp(ctx context.Context, email string, password string) (int64, error) {
	panic("not implemented")
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("not implemented")
}
