package auth

import (
	"my-rest-api/internal/domain/auth"
)

type Handler struct {
	Service auth.Service
}

func New(svc auth.Service) *Handler {
	return &Handler{
		Service: svc,
	}
}
