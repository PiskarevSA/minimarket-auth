package convtype

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var ErrTextIsNull = errors.New("int4 is null")

func TextToString(value pgtype.Text) (string, error) {
	if !value.Valid {
		return "", ErrTextIsNull
	}

	return value.String, nil
}

func StringToText(value string) pgtype.Text {
	result := pgtype.Text{String: value}

	if value == "" {
		return result
	}

	result.Valid = true

	return result
}
