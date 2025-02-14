package transfer

import (
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5"
)

type Transfer struct {
	queries schema.Queries
}

func New(conn *pgx.Conn) *Transfer {
	d := Transfer{
		queries: *schema.New(conn),
	}

	return &d
}
