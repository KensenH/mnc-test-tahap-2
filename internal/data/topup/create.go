package topup

import (
	"context"
	"my-rest-api/internal/schema"
)

func (d *TopUp) CreateTopUp(ctx context.Context, arg schema.CreateTopUpParams) (schema.TopUp, error) {
	topup, err := d.queries.CreateTopUp(ctx, arg)
	if err != nil {
		return topup, err
	}

	return topup, err
}
