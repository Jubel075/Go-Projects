/*
Copyright © 2025 Guilli guiliankasandikromo@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/jubel075/cli-cobra/todo"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your to-do list",
	Long: `The add command lets you create a new task and store it in your to-do list.

You can provide a short description of the task directly as an argument.
Optionally, you can attach metadata such as priority, due date, or tags
using flags. Tasks are saved in your local database or file storage so
they persist between sessions.

Examples:
  mytodo add "Buy groceries"
  mytodo add "Finish project report" --priority high --due tomorrow
  mytodo add "Call Alice" --tag personal --tag urgent

If no description is provided, the command will prompt you to enter one interactively.`,
	Run: func(cmd *cobra.Command, args []string) {
		var items = []todo.Item{}
		items, err := todo.ReadItems(dataFile)
		if err != nil {
			log.Printf("%v", err)
		}
		for _, x := range args {
			item := todo.Item{Text: x}
			item.SetPtiority(priority)
			items = append(items, item)
			fmt.Println("Added task:", item)
		}
		err = todo.SaveItems(dataFile, items)
		if err != nil {
			log.Printf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority of the task (1=high, 2=medium, 3=low)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
