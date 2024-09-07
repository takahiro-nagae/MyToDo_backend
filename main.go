package main

import (
	"MyToDo_backend/models"
)

func main() {
	// u, _ := models.GetUser(3)

	todo, _ := models.GetTodo(3)
	todo.DeleteTodo()
}
