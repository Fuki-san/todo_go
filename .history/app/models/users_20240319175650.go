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

//getuser関数は、レシーバを使わない。理由は、インスタンスに依存せずにデータベースに依存してデータを取得しているから。
//レシーバーを使うのは、createuserのようにインスタンスの構造に基づいて(依存して)データを作成するとき
func GetUser(id int) (user User, err error){
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error){
	cmd := `UPDATE users `
}

