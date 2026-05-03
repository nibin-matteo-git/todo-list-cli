package main

import (
	"todo-list-cli/internal/todo"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main(){
	todo.InitDB()
	log.Println("initdb done successfully")
	todo.GetTodos()
}

var DB *sql.DB
