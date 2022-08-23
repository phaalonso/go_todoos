/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/phaalonso/todoos/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new item to the todo list",
	Long:  `Add a new item to the todo list.`,
	Run:   addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(dataFile, items)

	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
