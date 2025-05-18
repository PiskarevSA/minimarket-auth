package usecases

import (
	"context"

	"github.com/google/uuid"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/entities"
	"github.com/PiskarevSA/minimarket-auth/internal/domain/objects"
)

type (
	accountRepo interface {
		CreateAccountInTx(ctx context.Context, account entities.Account) error
		GetUserIdAndPasswordHash(
			ctx context.Context,
			login objects.Login,
		) (uuid.UUID, []byte, error)
	}
)
