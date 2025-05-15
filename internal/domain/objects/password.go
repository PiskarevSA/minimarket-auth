package objects

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password []byte

var (
	ErrPasswordTooShort = errors.New("password too short")
	ErrPasswordTooLong  = errors.New("password too long")
)

var NilPassword Password = nil

func NewPassword(value string) (Password, error) {
	passwordLen := len(value)
	if passwordLen < 12 {
		return NilPassword, ErrPasswordTooShort
	}

	if passwordLen > 32 {
		return NilPassword, ErrPasswordTooLong
	}

	return Password([]byte(value)), nil
}

func (o Password) String() string {
	return string(o)
}

func (o Password) Hash() []byte {
	hash, _ := bcrypt.GenerateFromPassword(o, bcrypt.DefaultCost)
	return hash
}

func (o Password) IsMatchHash(hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, o)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
