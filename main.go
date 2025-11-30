package main

import (
	"fmt"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/types"
)

func main() {
	d, err := time.Parse("02-01-2006", "07-12-2025")
	if err != nil {
		return
	}

	todo := types.Todo{
		ID:       0,
		Text:     "Finish this damn staff",
		Done:     false,
		CreateAt: time.Now(),
		Due:      d,
	}

	todoList := db.TodoList{}
	todoList.Todos = append(todoList.Todos, todo)

	fmt.Println(todoList.Todos)
}
