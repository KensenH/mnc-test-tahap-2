package auth

import (
	"context"
	"my-rest-api/internal/domain/auth"
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s Auth) Register(ctx context.Context, req auth.RegisterReq) (schema.User, error) {
	return s.Data.CreateUser(ctx, schema.CreateUserParams{
		FirstName: req.FirstName,
		LastName: pgtype.Text{
			String: req.LastName,
			Valid:  true,
		},
		PhoneNumber: pgtype.Text{
			String: req.PhoneNumber,
			Valid:  true,
		},
		Address: pgtype.Text{
			String: req.PhoneNumber,
			Valid:  true,
		},
		Pin: req.Pin,
	})
}
