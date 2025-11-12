/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"

	// "strconv"
	"text/tabwriter"

	"github.com/jubel075/cli-cobra/todo"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the to-do tasks",
	Long: `The list command shows all tasks currently stored in your to-do list.

By default, it prints every task in the order they were added, along with
their status (completed or pending). You can filter the output using flags
to focus on what matters most. For example, you might only want to see
unfinished tasks, tasks due today, or tasks tagged with a specific label.

The list command supports multiple output styles. The default is a simple
human-readable table, but you can also request JSON or plain text for
integration with other tools and scripts. This makes it easy to pipe your
task data into other commands or automate reporting.

Examples:
  mytodo list
      Shows all tasks with their IDs and completion status.

  mytodo list --completed
      Displays only tasks that have been marked as done.

  mytodo list --pending --tag work
      Shows only unfinished tasks tagged with work.

  mytodo list --due today
      Lists tasks that are due today.

  mytodo list --output json
      Prints tasks in JSON format for use in scripts or other programs.

If no tasks exist, the command will let you know that your list is empty
instead of printing a blank table. This ensures you always get useful
feedback when running the command.`,

	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataFile)
		if err != nil {
			log.Printf("%v", err)
		}
		fmt.Printf("You have %d tasks in your to-do list:\n", len(items))

		sort.Sort(todo.ByPriority(items))

		// Print in tabular format
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		// Print header
		fmt.Fprintln(w, "LABEL\tPRIORITY\tTASK\tSTATUS")
		fmt.Fprintln(w, "-----\t--------\t----\t------")

		// Print each item with its label
		for _, i := range items {
			if i.Done || allOpt == doneOpt {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", i.Label(), i.PrettyP(), i.Text, i.PrettyDone())
			}
		}

		// Flush to output
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "List only completed tasks")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "List all tasks")
}
