package convtype

import (
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

var (
	ErrNumericIsNull     = errors.New("numeric is null")
	ErrNumericIsNan      = errors.New("numeric is nan")
	ErrNumericIsInfinite = errors.New("numeric is infinite")
)

func NumericToDecimal(value pgtype.Numeric) (decimal.Decimal, error) {
	if !value.Valid {
		return decimal.Decimal{}, ErrNumericIsNull
	}

	if value.NaN {
		return decimal.Decimal{}, ErrNumericIsNan
	}

	if value.InfinityModifier != pgtype.Finite {
		return decimal.Decimal{}, ErrNumericIsInfinite
	}

	return decimal.NewFromBigInt(value.Int, value.Exp), nil
}

func DecimalToNumeric(value decimal.Decimal) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   value.Coefficient(),
		Exp:   value.Exponent(),
		Valid: true,
	}
}
