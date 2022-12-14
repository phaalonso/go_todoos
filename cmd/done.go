/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/phaalonso/todoos/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark an task as done",
	Long:    `Select one or more task's to mark it as done`,
	Run:     markAsDone,
}

func markAsDone(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	items, err := todo.ReadItems(dataFile)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("File %s does not exist", dataFile)
		} else {
			log.Printf("%v", err)
		}
		return
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

		if num > 0 && num < len(items) {
			num -= 1
			items[num].Done = !items[num].Done
			fmt.Printf("%q %v\n", items[num].Text, "marked as done")
			sort.Sort(todo.ByPri(items))
		} else {
			log.Println(num, "doesn't match any items")
		}
	}

	todo.SaveItems(dataFile, items)
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
