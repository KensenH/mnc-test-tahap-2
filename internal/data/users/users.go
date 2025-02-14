package users

import (
	"my-rest-api/internal/schema"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	queries schema.Queries
}

func New(conn *pgx.Conn) *Users {
	d := Users{
		queries: *schema.New(conn),
	}

	return &d
}
