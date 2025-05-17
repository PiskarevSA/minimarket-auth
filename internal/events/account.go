package events

import (
	"time"

	"github.com/google/uuid"
)

type AccountRegisteredPayload struct {
	UserID uuid.UUID `json:"userId"`
	Login  string    `json:"login"`
}

const EventAccountRegistered = "ACCOUNTS.REGISTERED"

var AccountRegistered event

func NewAccountRegistered(
	userId uuid.UUID,
	login, createdBy string,
	createdAt time.Time,
) (Event, error) {
	payload := AccountRegisteredPayload{
		UserID: userId,
		Login:  login,
	}

	return NewEvent(EventAccountRegistered, payload, createdBy, createdAt)
}
