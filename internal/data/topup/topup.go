package topup

import (
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5"
)

type TopUp struct {
	queries schema.Queries
}

func New(conn *pgx.Conn) *TopUp {
	d := TopUp{
		queries: *schema.New(conn),
	}

	return &d
}
