package main

import (
	"MyToDo_backend/models"
	"fmt"
)

func main() {
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()

	u, _ = models.GetUser(1)
	fmt.Println(u)
}
