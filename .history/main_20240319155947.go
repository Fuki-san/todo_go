package main

import (
	"fmt"
	"todo_go/app/models"
)

func main(){
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test[]"
}
