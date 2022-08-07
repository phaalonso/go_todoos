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
	"strconv"
	"strings"

	"github.com/phaalonso/todoos/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark an task as done",
	Long:  `Select one or more task's to mark it as done`,
	Run:   markAsDone,
}

func markAsDone(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	items, err := todo.ReadItems("todoos.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("File %s does not exist", "todoos.json")
		} else {
			log.Printf("%v", err)
		}
		return
	}

	for i, x := range items {
		checked := " "
		if x.Done {
			checked = "✓"
		}
		fmt.Printf("%s %d - %s\n", checked, i+1, x.Text)
	}

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

		items[num-1].Done = !items[num-1].Done
	}

	todo.SaveItems("todoos.json", items)
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
