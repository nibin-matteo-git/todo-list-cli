package internal

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

// TODO: move table name to common config
var TABLE_NAME string = "todos"

type Action string

const (
	UPDATE Action = ""
	DELETE Action = ""
	CREATE Action = ""
)

type Event struct {
	gorm.Model
	Event        Action
	TodoId       uint
	TodoAffected Todo
}

type Todo struct {
	gorm.Model
	Name        string
	Description string
	Due         time.Time
	Done        bool
}

var DB *gorm.DB

var insertDummyData string = fmt.Sprintf(`
	insert into %s
	values (null, "test todo", "test-todo", "%v", "%v", 0, 0);`, TABLE_NAME, time.Now().Format(time.DateTime), time.Now().Format(time.DateTime))

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./todos.db"), &gorm.Config{})
	handleError(err, "Error opening sqlite connection")

	DB.AutoMigrate(
		&Todo{},
		&Event{},
	)
}
