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
  done <id>                Complete task by id
  delete <id>                Delete task by id
Examples:
  todo add "Buy groceries" "15-12-2025"
  todo update text "New text" done y 1
  todo delete 5
`
	fmt.Print(usage)
}

func printHorizontalRule(len int) {
	for i := 0; i < len; i++ {
		fmt.Print("-")
	}
	fmt.Println()
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
	parseFirstArgId := func() (int, error) {
		if len(args) == 0 {
			return 0, fmt.Errorf("Usage: %s <id>", action)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return 0, fmt.Errorf("invalid id '%s': %w", args[0], err)
		}

		return id, nil
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

		fmt.Print("\n\n")
		for i, t := range todos {

			status := "❌"
			if t.Done {
				status = "✅"
			}
			if i == 0 {
				printHorizontalRule(30)
			}
			fmt.Printf(
				"🆔 ID:      %d\n"+
					"✏️ Text:    %s\n"+
					"Status:    %s\n"+
					"🛠️ Created: %s\n"+
					"⏰ Due:     %s\n",
				t.ID,
				t.Text,
				status,
				t.CreateAt.Format("02-01-2006 15:04"),
				t.Due.Format("02-01-2006 15:04"),
			)
			printHorizontalRule(30)
		}
		fmt.Print("\n\n")

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
				db.UpdateTodo(todoList, id, map[string]any{
					"due": dueParsed,
				})
			case "text":
				db.UpdateTodo(todoList, id, map[string]any{
					"text": value,
				})

			case "done":
				db.UpdateTodo(todoList, id, map[string]any{
					"done": (value == "y" || value == "yes" || value == "true" || value == "1"),
				})
			default:
				fmt.Printf("Unknown field: %s\n", field)
			}
		}
		fmt.Printf("Successfully updated task!")
		if todo.Done {
			fmt.Printf("Todo with id %d is done", todo.ID)
		}
	case "done":
		id, err := parseFirstArgId()
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		db.UpdateTodo(todoList, id, map[string]any{
			"done": true,
		})
	case "clear":
		db.DeleteAll(todoList)
	case "delete":
		// Expect args: [id]
		id, err := parseFirstArgId()
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		if !db.DeleteTodo(todoList, id) {
			fmt.Println("Failed to delete todo. Probably the todo with such ID doesn't exist")
			return
		}
		fmt.Println("Todo deleted.")
	default:
		PrintUsage()
	}

	db.SaveTodos(todoList)
}
