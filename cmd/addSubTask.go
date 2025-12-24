/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/CristianHuang/cliTaskManager/internal/task"
	"github.com/spf13/cobra"
)

// addSubTaskCmd represents the addSubTask command
var addSubTaskCmd = &cobra.Command{
	Use:   "addSubTask",
	Short: "Add a subtask to a task",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if len(args) == 1 {
			fmt.Println("Please provide a subtask")
			return
		}

		list, err := task.Load()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		list.AddSubTask(id, args[1])
		err = task.Save(list)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Println("Subtask added:", args[1])
	},
}

func init() {
	rootCmd.AddCommand(addSubTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addSubTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addSubTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
