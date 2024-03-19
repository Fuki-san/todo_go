package models

import (
	"database/sql"
	"todo_go/config"
	_"github.com/mattn/go-sqlite3"
	"log"
	"fmt"
	"github.com/google/uuid"
	"crypto/sha1"
)

var Db *sql.DB

var err error

const(
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init(){
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil{
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`, tableNameUser)

	_, err := Db.Exec(cmdU);
	if err != nil{
		log.Fatalf("Fatled to create table %s : %v", tableNameUser, err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		user_id INTEGER,
		Cre
	)`, tableNameTodo)
}

func CreateUUID() (uuidobj uuid.UUID){
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
	}

func Encrypt(plaintext string) (cryptext string){
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
