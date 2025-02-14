package auth

import (
	"context"
	"my-rest-api/internal/schema"
)

type Service interface {
	Register(ctx context.Context, req RegisterReq) (schema.User, error)
	Login(ctx context.Context, req LoginReq) (string, error)
	VerifyToken(ctx context.Context, tokenString string) (string, error)
}
