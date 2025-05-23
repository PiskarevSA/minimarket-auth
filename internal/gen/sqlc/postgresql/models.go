// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package postgresql

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id           uuid.UUID
	Login        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Outbox struct {
	Id        int64
	Eventname string
	Status    string
	Payload   []byte
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
