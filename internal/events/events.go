package events

type Event string

func (e Event) String() string {
	return string(e)
}

const (
	AccountRegistered Event = "ACCOUNTS.REGISTERED"
)

type AccountRegisteredPayload struct {
	UserId string `json:"userId"`
	Login  string `json:"login"`
}
