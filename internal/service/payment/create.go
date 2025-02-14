package payment

import (
	"context"
	"math/big"
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Payment) Create(ctx context.Context, userId string, amount float64, remarks string) (schema.Payment, error) {
	var (
		err    error
		result schema.Payment
	)

	uuid := pgtype.UUID{}

	err = uuid.Scan(userId)
	if err != nil {
		return result, err
	}

	result, err = s.Data.CreatePayment(ctx, schema.CreatePaymentParams{
		Amount: pgtype.Numeric{
			Int:   big.NewInt(int64(amount)),
			Valid: true,
		},
		Remarks: remarks,
		UserID:  uuid,
	})

	//TO-DO: Create background process

	return result, err
}
