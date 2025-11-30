package types

import "time"

type Todo struct {
	ID       int
	Text     string
	Done     bool
	CreateAt time.Time
	Due      time.Time
}
type PartialTodo struct {
	Text string
	Done bool
	Due  time.Time
}

type TodoOptional struct {
	Text     *string
	Done     *bool
	Due      *time.Time
	CreateAt *time.Time
}
