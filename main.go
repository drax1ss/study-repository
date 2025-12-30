package main

import (
	"context"
	"fmt"
	simpleconection "obuchenie/feauture4/simple_conection"
	simplesql "obuchenie/feauture4/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconection.CheckConection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	// if err := simplesql.InsertRow(ctx,
	// 	conn,
	// 	"поиграть",
	// 	"Включить комп",
	// 	false,
	// 	time.Now()); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Успешно!")

	// if err := simplesql.UpdateRow(ctx, conn); err != nil {
	// 	panic(err)
	// }
	// if err := simplesql.DeleteSql(ctx, conn); err != nil {
	// 	panic(err)
	// }
	tasks, err := simplesql.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "dadad"
			task.Desription = "dadad"
			task.Completed = true
			task.Created_at = time.Now()

			if err := simplesql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
		}
		break
	}
	fmt.Print("succead")
}
