package models

import (
	"log"
	"time"
)

type User struct{
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

func (u *User) CreateUser() error{
	cmd := `INSERT INTO users (uuid, name, email, password, created_at) VALUES (?,?,?,?,?)`
	//first argument is command, other are values ?,?,?,?,?
	_, err := Db.Exec(cmd, CreateUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil{
		log.Fatalln(err)
	}
	return err
}

// func (u *User) GetUser() error {
// 	cmd := `SELECT id`
// 	_, err := Db.Exec(cmd, )
// }

func GetUser(id int) {user User, err error}{
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user
	)
}

