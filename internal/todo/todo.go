package todo

import (
	"time"
	"errors"
	"fmt"
)

type Todo struct {
	Name        string
	Description string
	Created     time.Time
	Due         time.Time
	Done		bool
}

func NewTodo(name string, description string, due time.Time) Todo{
	return Todo{
		Name: name,
		Description: description,
		Due: due,
		Created: time.Now(),
		Done: false,
	}
}

type List struct {
	Items []Todo
}

func NewList() *List{
	return &List{
		Items: []Todo{}, 
	}
}

func (l *List) AddTodo(text string, description string, due time.Time){
	newTodo := NewTodo(text, description, due)
	l.Items = append(l.Items, newTodo)
}

func (l *List) ToString() string{
	outStr := "Todos :\n"
	for i, v := range l.Items{
		status := " "
		if v.Done == true{
			status = "X"
		}
		outStrLine := fmt.Sprintf("%d. [%s] %s\n", i, status, v.Name)
		outStr += outStrLine
	}
	return outStr
}

func (l *List) CompleteTodo(index int) error{
	if (index < 0 || index > len(l.Items)){
		return errors.New("Index out of range")
	}
	l.Items[index].Done = true
	return nil
}
