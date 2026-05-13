package todo

import (
	"fmt"
	"log"
	"time"
)

type Todo struct {
	PKey        *int
	Name        string
	Description string
	Created     time.Time
	Due         time.Time
	Done        bool
	SyncStatus  bool
}

// TODO: write func to edit todo

func NewTodo(name string, description string, due time.Time) *[]Todo {
	insertIntoSql := fmt.Sprintf("INSERT INTO %s VALUES (null, %s, %s, %v, %v, 0, 0)", name, description, due.Format(time.DateTime), time.Now().Format(time.DateTime))
	_, err := DB.Exec(insertIntoSql)
	handleError(err, "Error adding new todo to DB")
	return GetTodos()
}

func CompleteTodo(pk int) *[]Todo{
	completeTodoSql := fmt.Sprintf("UPDATE %s set done = 1 where id = %d;", TABLE_NAME, pk)
	_, err := DB.Exec(completeTodoSql)
	handleError(err, "Unable to update todo as completed!!")
	return GetTodos()
}

func DeleteTodo(pk int) *[]Todo{
	deleteTodoSql := fmt.Sprintf("DELETE FROM %s WHERE id = %v;", TABLE_NAME, pk)
	_, err := DB.Exec(deleteTodoSql)
	handleError(err, "Unable to delete todo!!")
	return GetTodos()
}

func ( t *Todo ) PrintTodoDescription(){
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
	query := `Select * from ` + TABLE_NAME + `;`
	rows, err := DB.Query(query)
	handleError(err, "error querying select statement from db")
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.PKey, &todo.Name, &todo.Description, &todo.Created, &todo.Due, &todo.Done, &todo.SyncStatus)
		handleError(err, "Error when querying DB")
		result = append(result, todo)
	}

	return &result
}

func handleError(err error, msg string) {
	if err != nil {
		log.Println("Error message: " + msg)
		log.Println(err)
	}
}
