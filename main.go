package main

import (
	"MyToDo_backend/models"
)

func main() {
	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "text@example.com"
	u.Password = "testtest"
	u.CreateUser()
}
