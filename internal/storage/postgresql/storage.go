package postgresql

import (
	"github.com/PiskarevSA/minimarket-points/internal/gen/sqlc/postgresql"
)

type Storage struct {
	querier *postgresql.Queries
}

func New(dbtx postgresql.DBTX) *Storage {
	return &Storage{
		querier: postgresql.New(dbtx),
	}
}
