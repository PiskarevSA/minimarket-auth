package handlers

import (
	"context"

	"github.com/google/uuid"

	"github.com/PiskarevSA/minimarket-auth/internal/gen/oapi"
)

type (
	loginUsecase interface {
		Do(
			ctx context.Context,
			login string,
			password string,
		) (userId uuid.UUID, accessToken string, refreshToken string, err error)
	}

	registerUsecase interface {
		Do(
			ctx context.Context,
			login string,
			password string,
		) (userId uuid.UUID, accessToken string, refreshToken string, err error)
	}
)

var _ oapi.StrictServerInterface = (*handlers)(nil)

type handlers struct {
	loginUsecase    loginUsecase
	registerUsecase registerUsecase
}

func New(loginUsecase loginUsecase, registerUsecase registerUsecase) *handlers {
	return &handlers{
		loginUsecase:    loginUsecase,
		registerUsecase: registerUsecase,
	}
}
