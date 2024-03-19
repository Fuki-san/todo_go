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

func (u *User) GetUser() error {
	
}


