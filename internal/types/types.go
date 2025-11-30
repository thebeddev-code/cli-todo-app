package types

import "time"

type Todo struct {
	ID       int       `json:"id"`
	Text     string    `json:"text"`
	Done     bool      `json:"done"`
	CreateAt time.Time `json:"createAt"`
	Due      time.Time `json:"due"`
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
