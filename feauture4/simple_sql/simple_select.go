package simplesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRow(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id,title, desription,completed, created_at, completed_at 
	FROM tasks
	WHERE completed = TRUE
	ORDER BY id ASC
	 `

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Desription,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		task.PrintTask()
	}
	return tasks, nil
}

func (t TaskModel) PrintTask() {
	fmt.Println("______________________________________")
	fmt.Println("id:", t.ID)
	fmt.Println("title:", t.Title)
	fmt.Println("desription:", t.Desription)
	fmt.Println("completed:", t.Completed)
	fmt.Println("created_at:", t.Completed_at)
	fmt.Println("completed_at:", t.Completed_at)
}
