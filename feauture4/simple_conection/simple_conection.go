package simpleconection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:garfildR3@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err1 := conn.Ping(ctx); err1 != nil {
		panic(err1)
	}
	return conn, err

}
