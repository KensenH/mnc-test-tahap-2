package users

import (
	"context"
	"my-rest-api/internal/schema"
)

type Data interface {
	CreateUser(ctx context.Context, arg schema.CreateUserParams) (schema.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (schema.User, error)
}
