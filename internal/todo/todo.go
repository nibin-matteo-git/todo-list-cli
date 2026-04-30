package todo

import (
	"time"
	"errors"
	"fmt"
	"log"
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

 
func GetTodos() []Todo{
	var result []Todo
	query := `Select * from ` + TABLE_NAME + `;`
	rows, err := DB.Query(query)
	handleError(err, "error querying select statement from db")
	defer rows.Close()
	for rows.Next(){
		var todo Todo
		rows.Scan(&todo.Name, &todo.Description, &todo.Created, &todo.Due, &todo.Done)
		result = append(result, todo)
	}


	outStr := ""
	for i, v := range result{
		status := " "
		if v.Done == true{
			status = "X"
		}
		outStrLine := fmt.Sprintf("%d. [%s] %s\n", i, status, v.Name)
		outStr += outStrLine
	}
	fmt.Println(outStr)
	return result
}


func handleError(err error, msg string){
	if err != nil{
		log.Println(msg)
		log.Println(err)
	}
}
