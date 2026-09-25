package main

import (
	"log"
	"todo-list-cli/cmd"
	"todo-list-cli/internal"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	internal.InitDB()
	log.Println("initdb done successfully")
	internal.GetTodos()
	cmd.Execute()
}
