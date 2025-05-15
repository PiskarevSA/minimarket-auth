package signin

import (
	"context"

	"github.com/PiskarevSA/minimarket-points/internal/domain/entities"
	"github.com/PiskarevSA/minimarket-points/internal/domain/objects"
)

type Storage interface {
	GetAccountCredentials(
		ctx context.Context,
		login objects.Login,
	) (entities.Credentials, error)
}
