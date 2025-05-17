package objects

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password []byte

type PasswordError struct{ msg string }

func (e *PasswordError) Error() string {
	return e.msg
}

const (
	mixPasswordLen = 12
	maxPasswordLen = 24
)

var (
	ErrPasswordTooShort = &PasswordError{"password too short"}
	ErrPasswordTooLong  = &PasswordError{"password too long"}
)

var NilPassword Password = nil

func NewPassword(value string) (Password, error) {
	passwordLen := len(value)
	if passwordLen < mixPasswordLen {
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

func (o Password) IsHashMatch(hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, o)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
