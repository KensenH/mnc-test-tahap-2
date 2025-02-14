package topup

import (
	"my-rest-api/internal/domain/topup"
)

type TopUp struct {
	Data topup.Data
}

func New(data topup.Data) *TopUp {
	svc := TopUp{
		Data: data,
	}

	return &svc
}
