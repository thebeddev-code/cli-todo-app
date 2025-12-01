package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/types"
)

func PrintUsage() {
	const usage = `Usage: todo [command] [args]

Commands:
  add <text> [due-date]      Add new task (due-date: dd-mm-yyyy, default today)
  list                       List all tasks
  update <field> <value> [field value ...] <id>
                             Update task (fields: text, due, done)
  delete <id>                Delete task by id
  clear                      Clear screen
  q, quit                    Quit

Examples:
  todo add "Buy groceries" "15-12-2025"
  todo update text "New text" done y 1
  todo delete 5
`
	fmt.Print(usage)
}

// Reducer hah
func HandleAction(todoList *db.TodoList, action string, args []string) {
	const dateLayout = "02-01-2006"

	parseDate := func(input string) (time.Time, error) {
		if strings.TrimSpace(input) == "" {
			input = time.Now().Format(dateLayout)
		}
		return time.Parse(dateLayout, input)
	}

	switch action {
	case "add":
		// Expect args to be: [text, dueDate(optional)]
		if len(args) == 0 {
			fmt.Println("Text for new task is required.")
			return
		}
		text := args[0]
		dueStr := ""
		if len(args) > 1 {
			dueStr = args[1]
		}

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
		// Expect args: [field1, field2, ..., id, value1, value2, ...]
		// For simplicity: args format ["field", "value", "id"], support multiple fields by pairs + id at the end
		if len(args) < 3 || len(args)%2 == 0 {
			fmt.Println("Usage: update <field> <value> [<field> <value> ...] <id>")
			return
		}

		// The last argument is the id
		idStr := args[len(args)-1]
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

		// Process fields and values pairs before the id argument
		for i := 0; i < len(args)-1; i += 2 {
			field := args[i]
			value := args[i+1]

			switch field {
			case "due":
				dueParsed, err := parseDate(value)
				if err != nil {
					fmt.Println("Invalid date format, valid is: dd-mm-yyyy")
					return
				}
				todo.Due = dueParsed

			case "text":
				todo.Text = value

			case "done":
				val := strings.ToLower(strings.TrimSpace(value))
				todo.Done = (val == "y" || val == "yes" || val == "true" || val == "1")

			default:
				fmt.Printf("Unknown field: %s\n", field)
			}
		}
		fmt.Printf("Successfully updated task!")

	case "delete":
		// Expect args: [id]
		if len(args) != 1 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id")
			return
		}
		if !db.DeleteTodo(todoList, id) {
			fmt.Println("Failed to delete todo")
			return
		}
		fmt.Println("Todo deleted.")

	default:
		fmt.Println("Unknown command")
		PrintUsage()
	}

	db.SaveTodos(todoList)
}
