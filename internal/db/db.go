package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"todo-app/internal/types"
)

type Todo = types.Todo

type TodoList struct {
	Todos []Todo `json:"todos"`
}

func GetDataPath() string {
	// XDG_DATA_HOME (~/.local/share) or fallback to ~/.config
	if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
		return filepath.Join(dir, "todo", "todos.json")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "share", "todo", "todos.json")
}

func InitTodoList() (*TodoList, error) {
	path := GetDataPath()

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	// Load existing or create empty
	todos := &TodoList{}
	data, err := os.ReadFile(path)
	if err == nil {
		if err := json.Unmarshal(data, &todos.Todos); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return todos, nil
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
	for i, todo := range t.Todos {
		if todo.ID == id {
			return &t.Todos[i]
		}
	}
	return nil
}

func UpdateTodo(todoList *TodoList, id int, updates map[string]any) error {
	todo := GetTodo(todoList, id)
	if todo == nil {
		return fmt.Errorf("todo %d not found", id)
	}

	if text, ok := updates["text"].(string); ok {
		todo.Text = text
	}
	if done, ok := updates["done"].(bool); ok {
		todo.Done = done
	}
	if due, ok := updates["due"].(time.Time); ok {
		todo.Due = due
	}

	return nil
}

func DeleteTodo(t *TodoList, id int) bool {
	for i, todo := range t.Todos {
		if todo.ID == id {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			return true
		}
	}
	return false
}

func SaveTodos(t *TodoList) {
	data, err := json.MarshalIndent(t.Todos, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(GetDataPath(), data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
