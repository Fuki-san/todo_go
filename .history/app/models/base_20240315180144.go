package models

import "database/sql"

var Db *sql.DB

var err error

const{
	tableNameUser = "users"
}

func init(){
	Db, err = sql.Open(config.Config)
}
