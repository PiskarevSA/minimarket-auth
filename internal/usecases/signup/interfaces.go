package signup

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/PiskarevSA/minimarket-points/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-points/internal/events"
)

type Storage interface {
	CreateAccountInTx(
		ctx context.Context,
		userId uuid.UUID,
		login objects.Login,
		passwordHash []byte,
		createdAt time.Time,
	) error

	CreateOutboxInTx(
		ctx context.Context,
		event events.Event,
		payload []byte,
		createdAt time.Time,
		createdBy string,
	) error
}
