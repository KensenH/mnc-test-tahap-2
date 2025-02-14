package payment

import (
	"context"
	"my-rest-api/internal/schema"
)

func (d *Payment) CreatePayment(ctx context.Context, arg schema.CreatePaymentParams) (schema.Payment, error) {
	payment, err := d.queries.CreatePayment(ctx, arg)
	if err != nil {
		return payment, err
	}

	return payment, err
}
