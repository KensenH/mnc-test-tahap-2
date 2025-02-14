package payment

import (
	"my-rest-api/internal/domain/payment"
)

type Payment struct {
	Data payment.Data
}

func New(data payment.Data) *Payment {
	svc := Payment{
		Data: data,
	}

	return &svc
}
