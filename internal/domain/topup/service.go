package topup

import (
	"context"
	"my-rest-api/internal/schema"
)

type Service interface {
	Create(ctx context.Context, userId string, amount float64) (schema.TopUp, error)
}
