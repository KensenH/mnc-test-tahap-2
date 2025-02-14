package topup

import (
	"context"
	"math/big"
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *TopUp) Create(ctx context.Context, userId string, amount float64) (schema.TopUp, error) {
	var (
		err    error
		result schema.TopUp
	)

	uuid := pgtype.UUID{}

	err = uuid.Scan(userId)
	if err != nil {
		return result, err
	}

	result, err = s.Data.CreateTopUp(ctx, schema.CreateTopUpParams{
		AmountTopUp: pgtype.Numeric{
			Valid: true,
			Int:   big.NewInt(int64(amount)),
		},
		UserID: uuid,
	})

	//TO-DO: Create background process

	return result, err
}
