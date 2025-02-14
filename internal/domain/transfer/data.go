package transfer

import (
	"context"
	"my-rest-api/internal/schema"
)

type Data interface {
	CreateTransfer(ctx context.Context, arg schema.CreateTransferParams) (schema.Transfer, error)
}
