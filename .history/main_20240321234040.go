package main

import (
	"fmt"
	"todo_go/app/models"
)

func main(){
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@exmaple.com"
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
	// user, _ := models.GetUser(1)
	// user.CreateTodo("test todo")
	// t,_ := models.GetTodo(1)
	// fmt.Println(t)
	// user, _ := models.GetUser(2)
	// user.CreateTodo("test todo2")
	// todos, _ := models.GetTodosByUser()
	// for _, v := range todos{
	// 	fmt.Println(v)
	// }
	user, _ := models.GetUser(3)
	user.CreateTodo("Third todo")

	user, _ := models.GetUser(2)
	todos, _ := user.GetTodosByUser()
	for _, v := range todos{
		fmt.Println(v)
	}
	
}
