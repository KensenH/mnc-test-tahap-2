package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Connection *pgx.Conn
}

func New(ctx context.Context, user string, password string, dbName string, host string, port string) (Database, error) {
	var (
		err error
		db  Database
	)

	db.Connection, err = pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName))
	if err != nil {
		return db, err
	}

	return db, err
}

func (d *Database) Close(ctx context.Context) error {
	return d.Connection.Close(ctx)
}
