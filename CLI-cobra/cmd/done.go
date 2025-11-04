/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/jubel075/cli-cobra/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataFile)
		if err != nil {
			log.Printf("%v", err)
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task number:", args[0], err)
			return
		}
		if i > 0 && i < len(items) {
			items[i-1].Done = true
			fmt.Printf("%q %v\n", items[i-1].Text, "marked as done")
			sort.Sort(todo.ByPriority(items))
			todo.SaveItems(dataFile, items)
		} else {
			log.Fatalln("Task number out of range:", i)
		}
	},
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
