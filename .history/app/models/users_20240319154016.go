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
	cmd := "INSERT INTO user"
}
