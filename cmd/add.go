package cmd

import (
	"strings"
	"time"
	"todo-list-cli/internal/todo"

	"github.com/spf13/cobra"
)

var dueDate time.Time
var testVal string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `
	The allowed date time formats for the -t:
		DateTime   = "2006-01-02 15:04:05"
		DateOnly   = "2006-01-02"
		Kitchen     = "3:04PM"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _, _ := dueDate.Date()
		if (year == 0){
			dueDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), dueDate.Hour(), dueDate.Minute(), dueDate.Second(), 0, time.Now().Location())
		}

		var todos = todo.NewTodo(strings.Join(args, " "),"", dueDate)
		todo.PrintTodos(*todos)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	// TODO add timezones and set it as the machine's default timezone in DB
	addCmd.Flags().TimeVarP(&dueDate, "due", "t", time.Now().Add(time.Duration(96)*time.Hour), []string{time.Kitchen, time.DateTime, time.DateOnly},  "To describe due date")
}
