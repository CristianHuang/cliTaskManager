/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/CristianHuang/cliTaskManager/internal/task"
	"github.com/spf13/cobra"
)

const (
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	reset  = "\x1b[0m"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := task.Load()
		if err != nil {
			fmt.Println("Error loading tasks")
			return
		}
		if len(list) == 0 {
			fmt.Println("No tasks added")
			return
		}

		for _, t := range list {
			if t.Done {
				fmt.Printf("%s[x] %d %s%s\n", green, t.ID, t.Title, reset)
			} else {
				fmt.Printf("%s[ ] %d %s%s\n", yellow, t.ID, t.Title, reset)
			}
		}

	},
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
