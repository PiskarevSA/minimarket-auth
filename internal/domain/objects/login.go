package objects

import "errors"

type Login string

var (
	ErrLoginTooShort = errors.New("login too short")
	ErrLoginTooLong  = errors.New("login too long")
)

var NilLogin Login = ""

func NewLogin(value string) (Login, error) {
	loginLen := len(value)

	if loginLen < 5 {
		return NilLogin, ErrLoginTooShort
	}

	if loginLen > 24 {
		return NilLogin, ErrLoginTooLong
	}

	return Login(value), nil
}

func (o Login) String() string {
	return string(o)
}
