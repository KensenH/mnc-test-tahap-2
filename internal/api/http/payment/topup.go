package payment

import (
	"my-rest-api/internal/domain/payment"
)

type Handler struct {
	Service payment.Service
}

func New(svc payment.Service) *Handler {
	return &Handler{
		Service: svc,
	}
}
