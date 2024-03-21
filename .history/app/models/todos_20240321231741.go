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

func (u *User) CreateTodo(content string) (err error){
	cmd := `INSERT INTO todos (content, user_id, created_at) VALUES (?,?,?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo , err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = ?`
	err = Db.QueryRow(cmd, t.ID).Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)
	return todo, err
}

func (t *Todo) UpdateTodo() (err error){
	cmd := `UPDATE todos SET content = ? WHERE id = ?`
	_, err = Db.Exec(cmd, t.Content, t.ID)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (t *Todo) DeleteTodo() (err error){
	cmd := `DELETE FROM todos WHERE id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil{
		log.Fatal(err)
	}
	return err
}
