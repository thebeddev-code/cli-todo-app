package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/types"
)

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Reducer hah
func HandleAction(scanner *bufio.Scanner, todoList *db.TodoList, action string, shouldExit *bool) {
	const dateLayout = "02-01-2006"

	readLine := func(prompt string) string {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return ""
		}
		return scanner.Text()
	}

	parseDate := func(input string) (time.Time, error) {
		if strings.TrimSpace(input) == "" {
			input = time.Now().Format(dateLayout)
		}
		return time.Parse(dateLayout, input)
	}

	switch action {
	case "add":
		text := readLine("Enter text: ")
		dueStr := readLine("Enter due date (dd-mm-yyyy, default today): ")

		dueParsed, err := parseDate(dueStr)
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
		if len(todos) == 0 {
			fmt.Println("No tasks yet.")
			return
		}

		for _, t := range todos {
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
				t.CreateAt.Format("02-01-2006 15:04"),
				t.Due.Format("02-01-2006 15:04"),
			)
		}

	case "update":
		fmt.Print("Enter fields to update [text, due, done] and task id (e.g. \"text due 3\"): ")
		if !scanner.Scan() {
			return
		}
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Println("You must specify at least one field and an id.")
			return
		}

		idStr := parts[len(parts)-1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid id")
			return
		}

		todo := db.GetTodo(todoList, id)
		if todo == nil {
			fmt.Println("No todo with such id")
			return
		}

		fields := parts[:len(parts)-1]
		for _, field := range fields {
			switch field {
			case "due":
				dueStr := readLine("Enter new due date (dd-mm-yyyy, default today): ")
				dueParsed, err := parseDate(dueStr)
				if err != nil {
					fmt.Println("Invalid date format, valid is: dd-mm-yyyy")
					return
				}
				todo.Due = dueParsed

			case "text":
				text := readLine("Enter new text: ")
				todo.Text = text

			case "done":
				val := strings.TrimSpace(readLine("Enter y if done or n otherwise: "))
				todo.Done = (val == "y" || val == "Y")
			default:
				fmt.Printf("Unknown field: %s\n", field)
			}
		}

	case "delete":
		fmt.Print("Enter id to delete: ")
		if !scanner.Scan() {
			return
		}
		idStr := strings.TrimSpace(scanner.Text())
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid id")
			return
		}
		if db.DeleteTodo(todoList, id) == false {
			fmt.Println("Failed to delete todo:", err)
			return
		}
		fmt.Println("Todo deleted.")
	case "clear":
		ClearScreen()
	case "q", "quit":
		*shouldExit = true

	default:
		fmt.Println("Unknown action")
	}

	db.SaveTodos(todoList)
}
