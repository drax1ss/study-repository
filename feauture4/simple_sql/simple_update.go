package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateTask(ctx context.Context,
	conn *pgx.Conn,
	taskModel TaskModel) error {
	sqlQuery := `
	UPDATE tasks 
	SET title=$1, desription=$2, completed=$3, Completed_at=$4
	WHERE id=$5
	`

	if _, err := conn.Exec(ctx, sqlQuery, taskModel.Title,
		taskModel.Desription,
		taskModel.Completed,
		taskModel.Completed_at,
		taskModel.ID); err != nil {
		return err
	}
	return nil
}
