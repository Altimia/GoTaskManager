package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotask",
	Short: "GoTask is a command-line task manager",
	Long:  `A command-line task manager where users can add, view, update, delete tasks, and communicate with other users in real-time.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for running the command
	},
}

func Execute() {
	// code for executing the command
}
