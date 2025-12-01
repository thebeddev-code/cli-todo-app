package main

import (
	"log"
	"os"
	"todo-app/internal/db"
	"todo-app/internal/utils"
)

func main() {
	todoList, err := db.InitTodoList()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		utils.PrintUsage()
		return
	}

	action := os.Args[1]
	utils.HandleAction(todoList, action, os.Args[2:])
}
