package auth

import "my-rest-api/internal/domain/users"

type Auth struct {
	Data   users.Data
	Secret string
}

func New(data users.Data, secret string) *Auth {
	svc := Auth{
		Data:   data,
		Secret: secret,
	}

	return &svc
}
