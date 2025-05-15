package signup

import (
	"github.com/PiskarevSA/minimarket-points/internal/domain/objects"
	"github.com/PiskarevSA/minimarket-points/internal/usecases"
)

type args struct {
	Login    objects.Login
	Password objects.Password
}

var nilArgs = args{}

func newArgs(
	rawLogin string,
	rawPassword string,
) (args, error) {
	login, err := objects.NewLogin(rawLogin)
	if err != nil {
		return nilArgs,
			&usecases.ValidationError{
				Field:   "login",
				Message: err.Error(),
			}
	}

	password, err := objects.NewPassword(rawPassword)
	if err != nil {
		return nilArgs,
			&usecases.ValidationError{
				Field:   "password",
				Message: err.Error(),
			}
	}

	return args{
		Login:    login,
		Password: password,
	}, nil
}
