package models

import "time"

type Todo struct{
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

func (t *Todo) CreateTodo() error{
	cmd := ``
}
