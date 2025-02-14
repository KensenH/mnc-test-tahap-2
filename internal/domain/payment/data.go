package payment

import (
	"context"
	"my-rest-api/internal/schema"
)

type Data interface {
	CreatePayment(ctx context.Context, arg schema.CreatePaymentParams) (schema.Payment, error)
}
