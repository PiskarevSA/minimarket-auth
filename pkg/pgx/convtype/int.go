package convtype

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

var ErrIn4IsNull = errors.New("int4 is null")

func Int32ToInt4(value int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: value,
		Valid: true,
	}
}

func Int4ToInt32(value pgtype.Int4) (int32, error) {
	if !value.Valid {
		return 0, ErrIn4IsNull
	}

	return value.Int32, nil
}
