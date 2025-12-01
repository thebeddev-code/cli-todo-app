package main

import (
	"bufio"
	"fmt"
	"os"
	"todo-app/internal/db"
	"todo-app/internal/utils"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	todoList := db.NewTodoList()
	db.InitTodoList(todoList)
	shouldExit := false
	for !shouldExit {
		fmt.Print(Green + "\nChoose one action from the list: add | list | list one | update | delete -> " + Reset)
		var action string
		fmt.Scanf("%s", &action)
		utils.HandleAction(scanner, todoList, action, &shouldExit)
	}
}
