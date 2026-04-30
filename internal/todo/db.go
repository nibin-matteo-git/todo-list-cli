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
	values (4, "test todo", "test-todo", "%v", "%v", 0);`, TABLE_NAME, time.Now().Format(time.DateTime), time.Now().Format(time.DateTime))


var dbCreationSql string = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255), 
	description TEXT, 
	created DATETIME, 
	due DATETIME, 
	done BOOLEAN DEFAULT 0
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


