package main

import (
	"flag"
	"fmt"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/types"
)

func main() {
	text := flag.String("text", "default", "")
	today := time.Now().Format("02-01-2006")
	due := flag.String("due", today, "Due date (dd-mm-yyyy)")
	flag.Parse()
	dueParsed, err := time.Parse("02-01-2006", *due)
	if err != nil {
		fmt.Println("Invalid date format, valid is: dd-mm-yyyy")
		return
	}

	todoList := db.TodoList{}
	db.AddTodo(&todoList, types.PartialTodo{
		Text: *text,
		Due:  dueParsed,
		Done: false,
	})

	fmt.Println(todoList.Todos[0].Due)
}
