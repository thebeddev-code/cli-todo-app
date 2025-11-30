package db

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"todo-app/internal/types"
)

type Todo = types.Todo

type TodoList struct {
	Todos []Todo `json:"todos"`
}

func NewTodoList() *TodoList {
	todoList := TodoList{}
	todoList.Todos = []Todo{}
	return &todoList
}

func InitTodoList(t *TodoList) {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		t.Todos = []Todo{}
		return
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		log.Printf("Failed to parse todos.json: %v", err)
		t.Todos = []Todo{}
		return
	}
}

func GetUniqueId(t *TodoList) int {
	todos := GetTodos(t)
	if len(todos) == 0 {
		return 1
	}
	return todos[len(todos)-1].ID + 1
}

func AddTodo(t *TodoList, todo types.PartialTodo) {
	t.Todos = append(t.Todos, Todo{
		ID:       GetUniqueId(t),
		Text:     todo.Text,
		Done:     todo.Done,
		CreateAt: time.Now(),
		Due:      todo.Due,
	})
}

func GetTodos(t *TodoList) []Todo {
	return t.Todos
}

func GetTodo(t *TodoList, id int) *Todo {
	for i := range t.Todos {
		if t.Todos[i].ID == id {
			return &t.Todos[i]
		}
	}
	return nil
}

func UpdateTodo(t *TodoList, id int, todo *types.TodoOptional) *Todo {
	for i := range t.Todos {
		if t.Todos[i].ID == id {
			return &t.Todos[i]
		}
	}
	return nil
}

func SaveTodos(t *TodoList) {
	data, err := json.MarshalIndent(t, "", "  ") // Serialize TodoList to JSON
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("todos.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
