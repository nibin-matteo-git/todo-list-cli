package internal

import (
	"fmt"
	"log"
	"time"
)

// TODO: write func to edit todo

func NewTodo(name string, description string, due time.Time) *[]Todo {
	newTodo := Todo{Name: name, Description: description, Due: due, Done: false}
	DB.Create(&newTodo)
	return GetTodos()
}

func CompleteTodo(todo *Todo) *[]Todo {
	todo.Done = true
	DB.Save(todo)
	return GetTodos()
}

func DeleteTodo(todo *Todo) *[]Todo {
	DB.Delete(todo)
	return GetTodos()
}

func (t *Todo) PrintTodoDescription() {
	fmt.Println("Description: \n" + t.Description)
	fmt.Println("\nDue: \n" + t.Due.Format(time.DateTime))
}

func PrintTodos(todoList []Todo) {

	outStr := ""
	for i, v := range todoList {
		status := " "
		if v.Done == true {
			status = "X"
		}
		outStrLine := fmt.Sprintf("%d. [%s] %s\n", i, status, v.Name)
		outStr += outStrLine
	}
	fmt.Println(outStr)
}

func GetTodos() *[]Todo {
	var result []Todo
	DB.Find(&result)
	return &result
}

func handleError(err error, msg string) {
	if err != nil {
		log.Println("Error message: " + msg)
		log.Println(err)
	}
}
