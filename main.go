package main

import (
	"os"
	"todo-app/internal/db"
	"todo-app/internal/utils"
)

func main() {
	todoList := db.NewTodoList()
	db.InitTodoList(todoList)
	if len(os.Args) < 2 {
		utils.PrintUsage()
		return
	}

	action := os.Args[1]
	utils.HandleAction(todoList, action, os.Args[2:])
}
