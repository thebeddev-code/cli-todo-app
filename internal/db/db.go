package db

import (
	"time"
	"todo-app/internal/types"
)

type Todo = types.Todo

type TodoList struct {
	Todos []Todo
}

func GetUniqueId(t *TodoList) int {
	todos := GetTodos(t)
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
