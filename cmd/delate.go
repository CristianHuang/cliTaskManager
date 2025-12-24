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

// delateCmd represents the delate command
var delateCmd = &cobra.Command{
	Use:   "delate",
	Short: "Delete a task from the list",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		list, _ := task.Load()
		if !list.Delete(id) {
			fmt.Println("Task not found")
			return
		}
		task.Save(list)
		fmt.Println("Task deleted")
	},
}

func init() {
	rootCmd.AddCommand(delateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
