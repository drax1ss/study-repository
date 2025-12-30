package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQuery := `
	INSERT INTO tasks (title, desription, completed, created_at)
	VALUES($1, $2, $3, $4)
	`
	_, err := conn.Exec(ctx, sqlQuery,
		task.Title,
		task.Desription,
		task.Completed,
		task.Created_at)

	return err
}
