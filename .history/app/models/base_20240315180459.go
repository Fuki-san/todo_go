package models

import (
	"database/sql"
	"todo_go/config"
	_"github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const{
	tableNameUser = "users"
}

func init(){
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil{
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf('CREATE TABLE IF NOT EXITSTS %s(
		i
	)')
}
