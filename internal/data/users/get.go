package users

import (
	"context"
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5/pgtype"
)

func (d *Users) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (schema.User, error) {
	user, err := d.queries.GetUserByPhoneNumber(ctx, pgtype.Text{String: phoneNumber, Valid: true})
	if err != nil {
		return user, err
	}

	return user, err
}
