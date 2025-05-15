package entities

import "github.com/google/uuid"

type Credentials struct {
	UserId       uuid.UUID
	PasswordHash []byte
}

var NilCredentials = Credentials{}
