package events

import (
	"time"

	json "github.com/bytedance/sonic"
)

type event struct {
	Name      string
	Status    string
	Payload   []byte
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

const (
	StatusUnprocessed = "UNPROCESSED"
	StatusProcessing  = "PROCESSING"
	StatusProcessed   = "PROCESSED"
	StatusFailed      = "FAILED"
	StatusRetry       = "RETRY"
)

func NewEvent[PayloadT any](
	name string,
	payload PayloadT,
	createdBy string,
	createdAt time.Time,
) (Event, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &event{
		Name:      name,
		Status:    StatusUnprocessed,
		Payload:   data,
		CreatedAt: createdAt,
		CreatedBy: createdBy,
		UpdatedAt: createdAt,
		UpdatedBy: createdBy,
	}, nil
}

func (e *event) GetName() string {
	return string(e.Name)
}

func (e *event) GetPayload() []byte {
	return e.Payload
}

func (e *event) GetStatus() string {
	return e.Status
}

func (e *event) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *event) GetCreatedBy() string {
	return e.CreatedBy
}

func (e *event) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *event) GetUpdatedBy() string {
	return e.UpdatedBy
}
