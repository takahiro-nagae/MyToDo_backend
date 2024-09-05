package main

import (
	"MyToDo_backend/models"
	"fmt"
)

func main() {
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
