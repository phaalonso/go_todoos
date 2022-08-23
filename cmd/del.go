/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/phaalonso/todoos/todo"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an specific task item",
	Long:  `Use this command to list, and select an specif task item to be removed from your tasks`,
	Run:   deleteTask,
}

func deleteTask(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Printf("%v", err)
	}

	todo.ListItems(items)

	fmt.Println("Select one or more options to mark as done (separate with ',')")
	fmt.Print("-> ")

	text, err := reader.ReadString('\n')

	if len(text) == 0 || err != nil {
		return
	}

	values := strings.Split(text, ",")

	for _, x := range values {
		x = strings.TrimSpace(x)

		num, err := strconv.Atoi(x)

		if err != nil {
			log.Printf("%v", err)
			continue
		}

		if num <= 0 {
			log.Println("Reveiced number that is equals or lower than zero, proceding to ignore it")
			continue
		}

		num -= 1

		copy(items[num:], items[num+1:])
		items = items[:len(items)-1]
	}

	todo.SaveItems(dataFile, items)
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
