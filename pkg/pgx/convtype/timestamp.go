package convtype

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var ErrTimestamptzIsNull = errors.New("timestamptz is null")

func TimestamptzToTime(ts pgtype.Timestamptz) (time.Time, error) {
	if !ts.Valid {
		return ts.Time, ErrTimestamptzIsNull
	}

	if ts.InfinityModifier == pgtype.Infinity {
		return time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC), nil
	}

	if ts.InfinityModifier == pgtype.NegativeInfinity {
		return time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), nil
	}

	return ts.Time, nil
}
