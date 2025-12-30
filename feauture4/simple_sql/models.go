package simplesql

import "time"

type TaskModel struct {
	ID           int
	Title        string
	Desription   string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}
