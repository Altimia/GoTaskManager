package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotask",
	Short: "GoTask is a command-line task manager",
	Long:  `A command-line task manager where users can add, view, update, delete tasks, and communicate with other users in real-time.`,
}

func Execute() {
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(chatCmd)
	rootCmd.Execute()
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
	Long:  `Add, view, update, and delete tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for running the task command
	},
}

taskCmd.AddCommand(addTaskCmd, viewTaskCmd, updateTaskCmd, deleteTaskCmd)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Register, login, and manage user profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for running the user command
	},
}

userCmd.AddCommand(registerCmd, loginCmd, manageProfileCmd)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with users",
	Long:  `Send and receive chat messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for running the chat command
	},
}

chatCmd.AddCommand(sendMessageCmd, receiveMessageCmd)
