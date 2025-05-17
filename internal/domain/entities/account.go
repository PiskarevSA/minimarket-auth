package entities

import (
	"time"

	"github.com/google/uuid"

	"github.com/PiskarevSA/minimarket-auth/internal/domain/objects"
)

type Account struct {
	id        uuid.UUID
	login     objects.Login
	password  objects.Password
	createdAt time.Time
	updatedAt time.Time
}

var NilAccount = Account{}

func NewAccount(
	id uuid.UUID,
	login,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) (Account, error) {
	var (
		account Account
		err     error
	)

	account.login, err = objects.NewLogin(login)
	if err != nil {
		return NilAccount, err
	}

	account.password, err = objects.NewPassword(password)
	if err != nil {
		return NilAccount, err
	}

	account.id = id
	account.createdAt = createdAt
	account.updatedAt = updatedAt

	return account, nil
}

func (a *Account) Id() uuid.UUID {
	return a.id
}

func (a *Account) Login() objects.Login {
	return a.login
}

func (a *Account) Password() objects.Password {
	return a.password
}

func (a *Account) PasswordHash() []byte {
	return a.password.Hash()
}

func (a *Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Account) UpdatedAt() time.Time {
	return a.updatedAt
}
