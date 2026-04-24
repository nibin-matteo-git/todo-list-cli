package todo

import "time"

type Todo struct {
	Name        string
	Description string
	Created     time.Time
	Due         time.Time
}

func NewTodo(name string, description string, due time.Time) Todo{
	return Todo{
		Name: name,
		Description: description,
		Due: due,
		Created: time.Now(),
	}
}

