package handlers

import (
	"context"
	"errors"

	"github.com/PiskarevSA/minimarket-auth/internal/gen/oapi"
	"github.com/PiskarevSA/minimarket-auth/internal/usecases"
)

var login500JSONResponse = oapi.Login500JSONResponse{
	Code:    oapi.S1394,
	Message: oapi.InternalError,
}

func (h *handlers) Login(
	ctx context.Context,
	request oapi.LoginRequestObject,
) (oapi.LoginResponseObject, error) {
	userId, accessToken, refreshTokem, err := h.loginUsecase.Do(
		ctx,
		request.Body.Login,
		request.Body.Password,
	)
	if err != nil {
		var validationErr *usecases.ValidationError
		if errors.As(err, &validationErr) {
			return oapi.Login400JSONResponse{
				Code:    oapi.ValidationErrorCode(validationErr.Code),
				Field:   validationErr.Field,
				Message: validationErr.Message,
			}, nil
		}

		var businessErr *usecases.BusinessError
		if errors.As(err, &businessErr) {
			return oapi.Login422JSONResponse{
				Code:    oapi.DomainErrorCode(businessErr.Code),
				Message: businessErr.Message,
			}, nil
		}

		return login500JSONResponse, nil
	}

	return oapi.Login200JSONResponse{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshTokem,
	}, nil
}
