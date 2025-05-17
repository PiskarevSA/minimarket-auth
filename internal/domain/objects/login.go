package objects

type Login string

type LoginError struct{ msg string }

func (e *LoginError) Error() string {
	return e.msg
}

const (
	minLoginLen = 5
	maxLoginLen = 24
)

var (
	ErrLoginTooShort = &LoginError{"login too short"}
	ErrLoginTooLong  = &LoginError{"login too long"}
)

var NilLogin Login = ""

func NewLogin(value string) (Login, error) {
	loginLen := len(value)

	if loginLen < minLoginLen {
		return NilLogin, ErrLoginTooShort
	}

	if loginLen > maxLoginLen {
		return NilLogin, ErrLoginTooLong
	}

	return Login(value), nil
}

func (o Login) String() string {
	return string(o)
}
