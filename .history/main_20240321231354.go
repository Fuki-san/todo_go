package main

import (
	"fmt"
	"todo_go/app/models"
)

func main(){
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@exmaple.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()
	// u,_ := models.GetUser(1)
	// fmt.Println(u)
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
	user, _ := models.GetUser(2)
	user.

}
