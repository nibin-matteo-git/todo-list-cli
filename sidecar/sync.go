package sidecar

import (
	"todo-list-cli/internal"
)

type Action string

const (
	UPDATE Action = ""
	DELETE Action = ""
	CREATE Action = ""
)

type Event struct {
	PKey   *int
	Event  Action
	TodoId internal.Todo
}
