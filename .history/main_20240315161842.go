package main

import (
	"todo_go/config/config.go"
	"fmt"
)

func main(){
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
}
