package users

import (
	"context"
	"my-rest-api/internal/schema"
)

func (d *Users) CreateUser(ctx context.Context, arg schema.CreateUserParams) (schema.User, error) {
	user, err := d.queries.CreateUser(ctx, arg)
	if err != nil {
		return user, err
	}

	return user, err
}
