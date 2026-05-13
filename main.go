package main

import (
	"todo-list-cli/internal/todo"
	"todo-list-cli/cmd"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main(){
	todo.InitDB()
	log.Println("initdb done successfully")
	todo.GetTodos()
	cmd.Execute()
}
