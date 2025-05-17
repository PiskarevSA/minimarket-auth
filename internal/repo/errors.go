package repo

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

var (
	ErrLoginAlreadyInUse = &Error{"login already in use"}
	ErrLoginNotExists    = &Error{"login not exists"}
)
