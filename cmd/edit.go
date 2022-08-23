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

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an task item",
	Long:  `Edit task item description`,
	Run:   editTask,
}

func editTask(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Printf("%v", err)
	}

	todo.ListItems(items)

	text, err := reader.ReadString('\n')

	if len(text) == 0 || err != nil {
		return
	}

	fmt.Println("Digite o novo conteudo da tarefa:")
	fmt.Print("-> ")

	num, err := strconv.Atoi(strings.TrimSpace(text))

	if num <= 0 {
		log.Printf("Number received is lower or equals than zero, proceding to ignore it")
		return
	}

	text, err = reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error reading from stding")
	}

	items[num-1].Text = strings.TrimSpace(text)

	todo.SaveItems(dataFile, items)
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
