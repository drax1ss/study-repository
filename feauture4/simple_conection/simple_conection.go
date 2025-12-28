package simpleconection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключение успешно!!")
}
