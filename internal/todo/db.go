package todo

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var TABLE_NAME string = "todos"

var DB *sql.DB

var insertDummyData string = fmt.Sprintf(`
	insert into %s
	values (null, "test todo", "test-todo", "%v", "%v", 0, 0);`, TABLE_NAME, time.Now().Format(time.DateTime), time.Now().Format(time.DateTime))


var dbCreationSql string = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	Name VARCHAR(255), 
	Description TEXT, 
	Created DATETIME, 
	Due DATETIME, 
	Done BOOLEAN DEFAULT 0, 
	SyncStatus BOOLEAN DEFAULT 0
	);`, TABLE_NAME)


func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3", "./todos.db")
	handleError(err, "Error opening sqlite connection")

	_, err = DB.Exec(dbCreationSql)
	handleError(err, "error creating table")

	// _, err = DB.Exec(insertDummyData)
	// handleError(err, "error inserting dummy data")
}


