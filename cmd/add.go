/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var dueDate time.Time
var testVal string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		fmt.Println(args)
		// fmt.Printf("length %d", len(args))
		// fmt.Println("the timne is: ", dueDate)
		// fmt.Println("time after formatting: ", dueDate.Format(time.DateTime))

		fmt.Println(testVal)
		fmt.Println(time.Parse(time.DateOnly, testVal))
		fmt.Println(time.Parse(time.DateTime, testVal))
		fmt.Println(time.Parse(time.Kitchen, testVal))
	},
}

func init() {
	fmt.Println("adding root cmd now")
	rootCmd.AddCommand(addCmd)

	fmt.Println("adding flags now")

	// TODO add timezones and set it as the machine's default timezone in DB

	addCmd.Flags().TimeVarP(&dueDate, "due", "t", time.Now().Add(time.Duration(6)*time.Hour), []string{time.Kitchen, time.DateTime, time.DateOnly},  "To describe due date")

	addCmd.Flags().StringVar(&testVal, "test", "test-val", "test-val")
}
