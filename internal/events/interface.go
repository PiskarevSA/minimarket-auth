package events

import "time"

type Event interface {
	GetName() string
	GetPayload() []byte
	GetStatus() string
	GetCreatedAt() time.Time
	GetCreatedBy() string
	GetUpdatedAt() time.Time
	GetUpdatedBy() string
}
