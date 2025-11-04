/*
Copyright © 2025 Guilli guiliankasandikromo@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
		for _, x := range args {
			fmt.Println("Added task:", x)
		}
	},
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
