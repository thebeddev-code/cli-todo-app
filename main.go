package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/types"
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

	for {
		fmt.Print(Green + "\nChoose one action from the list: add | list | list one | update | delete -> " + Reset)
		var action string
		fmt.Scanf("%s", &action)
		switch strings.TrimSpace(action) {
		case "add":
			{
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
				break
			}
		case "list":
			{
				fmt.Println(db.GetTodos(todoList))
				break
			}
		case "list one":
			{
				break
			}
		case "update":
			{
				break
			}
		case "delete":
			{
				break
			}
		default:
			return
		}
	}
}
