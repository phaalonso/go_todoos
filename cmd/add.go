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

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems("todoos.json")

	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}

	err = todo.SaveItems("todoos.json", items)

	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
