/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/phaalonso/todoos/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the saved todoos",
	Long:  `List the saved todos`,
	Run:   listTodoos,
}

func listTodoos(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems("todoos.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("File %s does not exist", "todoos.json")
		} else {
			log.Printf("%v", err)
		}
	}

	todo.ListItems(items)
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
