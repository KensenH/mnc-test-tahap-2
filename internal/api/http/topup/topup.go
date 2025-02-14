package topup

import (
	"my-rest-api/internal/domain/topup"
)

type Handler struct {
	Service topup.Service
}

func New(svc topup.Service) *Handler {
	return &Handler{
		Service: svc,
	}
}
