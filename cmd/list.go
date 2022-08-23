/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/phaalonso/todoos/todo"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the saved todoos",
	Long:  `List the saved todos`,
	Run:   listTodoos,
}

func listTodoos(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("File %s does not exist", dataFile)
		} else {
			log.Printf("%v", err)
		}
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, "\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
		}
	}

	w.Flush()

	// todo.ListItems(items)
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show alll Todos")
}
