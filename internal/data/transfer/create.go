package transfer

import (
	"context"
	"my-rest-api/internal/schema"
)

func (d *Transfer) CreateTransfer(ctx context.Context, arg schema.CreateTransferParams) (schema.Transfer, error) {
	transfer, err := d.queries.CreateTransfer(ctx, arg)
	if err != nil {
		return transfer, err
	}

	return transfer, err
}
