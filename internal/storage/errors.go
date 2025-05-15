package storage

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

var (
	ErrLoginAlreadyExists         = &Error{"login already exists"}
	ErrAccountCredentialsNotFound = &Error{"account credentials not found"}
)
