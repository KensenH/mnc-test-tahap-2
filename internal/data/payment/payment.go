package payment

import (
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5"
)

type Payment struct {
	queries schema.Queries
}

func New(conn *pgx.Conn) *Payment {
	d := Payment{
		queries: *schema.New(conn),
	}

	return &d
}
