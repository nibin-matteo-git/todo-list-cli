package main

import (
	"fmt"
	"todo-list-cli/internal/todo"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main(){
	myList := todo.NewList()
	myList.AddTodo("trial", "", time.Now())
	fmt.Println(myList.ToString())
	myList.CompleteTodo(0)
	fmt.Println(myList.ToString())
	initDB()


}

var DB *sql.DB

func initDB(){
	var err error
	DB, err = sql.Open("sqlite3", "./todos.db")
	if err != nil{
		log.Println(err)
	}
	// sqlStmt := `
	// CREATE TABLE IF NOT EXISTS todos (
	// id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// title TEXT
	// );`
	sqlStmt := 	`select * from todos;`
	// _, err := DB.Exec(sqlStmt)
	// if err != nil {
	// 	  log.Fatal("Error creating table: %q: %s\n", err, sqlStmt) 
	// }

	rows, err := DB.Query(sqlStmt)
	if err != nil {
		  log.Fatal("Error creating table: %q: %s\n", err, sqlStmt) 
	}
	defer rows.Close()

}
