package topup

import (
	"context"
	"my-rest-api/internal/schema"
)

type Data interface {
	CreateTopUp(ctx context.Context, arg schema.CreateTopUpParams) (schema.TopUp, error)
}
