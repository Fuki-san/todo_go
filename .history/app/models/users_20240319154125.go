package models

import "time"

type User struct{
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

func (u *User) CreateUser() error{
	cmd := `INSERT INTO users (uuid, name, email, password, created_at) VALUSE (?,?,?,?,?)`
	_, err := Db.Exec()
}
