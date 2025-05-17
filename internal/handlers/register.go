package handlers

import (
	"context"
	"errors"

	"github.com/PiskarevSA/minimarket-auth/internal/gen/oapi"
	"github.com/PiskarevSA/minimarket-auth/internal/usecases"
)

var register500JSONResponse = oapi.Register500JSONResponse{
	Code:    oapi.S1394,
	Message: oapi.InternalError,
}

func (h *handlers) Register(
	ctx context.Context,
	request oapi.RegisterRequestObject,
) (oapi.RegisterResponseObject, error) {
	userId, accessToken, refreshTokem, err := h.registerUsecase.Do(
		ctx,
		request.Body.Login,
		request.Body.Password,
	)
	if err != nil {
		var validationErr *usecases.ValidationError
		if errors.As(err, &validationErr) {
			return oapi.Register400JSONResponse{
				Code:    oapi.ValidationErrorCode(validationErr.Code),
				Field:   validationErr.Field,
				Message: validationErr.Message,
			}, nil
		}

		var businessErr *usecases.BusinessError
		if errors.As(err, &businessErr) {
			return oapi.Register422JSONResponse{
				Code:    oapi.DomainErrorCode(businessErr.Code),
				Message: businessErr.Message,
			}, nil
		}

		return register500JSONResponse, nil
	}

	return oapi.Register200JSONResponse{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshTokem,
	}, nil
}
