package models

import (
	"log"
	"time"
)

type Todo struct{
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

func (t *Todo) CreateTodo() error{
	cmd := `INSERT INTO todos (content, user_id, created_at) VALUES (?,?,?)`
	_, err := Db.Exec(cmd, t.Content, t.UserID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) GetTodo() error {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = ?`
	
}
