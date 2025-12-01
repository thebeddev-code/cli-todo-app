package utils

import (
	"bufio"
	"fmt"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/types"
)

func HandleAction(scanner *bufio.Scanner, todoList *db.TodoList, action string, shouldExit *bool) {
	switch action {
	case "add":
		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()

		fmt.Print("Enter due date (dd-mm-yyyy, default today): ")
		scanner.Scan()
		due := scanner.Text()
		if due == "" {
			due = time.Now().Format("02-01-2006")
		}

		dueParsed, err := time.Parse("02-01-2006", due)
		if err != nil {
			fmt.Println("Invalid date format, valid is: dd-mm-yyyy")
			return
		}
		db.AddTodo(todoList, types.PartialTodo{
			Text: text,
			Done: false,
			Due:  dueParsed,
		})
		fmt.Println("New task added!")
	case "list":
		todos := db.GetTodos(todoList)
		todosLen := len(todos)
		if todosLen == 0 {
			return
		}
		for i := range todosLen {
			t := todos[i]
			status := "❌"
			if t.Done {
				status = "✅"
			}
			fmt.Printf(
				"🆔 ID: %d\n✏️  Text: %s\n%s Done: %v\n🛠️  Created: %s\n⏰ Due: %s\n\n",
				t.ID,
				t.Text,
				status,
				t.Done,
				t.CreateAt.Format("01-02-2006 15:04"),
				t.Due.Format("01-02-2006 15:04"),
			)
		}
	case "list one":
		break
	case "update":
		break
	case "delete":
		break
	case "q":
	case "quit":
		*shouldExit = true
	}
}
