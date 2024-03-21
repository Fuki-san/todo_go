package main

import (
	"fmt"
	"todo_go/app/controllers"
	"todo_go/app/models"
)

func main(){
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
