package main

import (
	"fmt"
	"strconv"

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
}

func init() {
	taskCmd.AddCommand(addTaskCmd)
	taskCmd.AddCommand(viewTaskCmd)
	taskCmd.AddCommand(updateTaskCmd)
	taskCmd.AddCommand(deleteTaskCmd)
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long:  `Add a new task to the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for adding a task
		if len(args) != 4 {
			fmt.Println("Invalid number of arguments. Expected 4 arguments: name, description, status, assignedTo.")
			return
		}
		var user User
		if err := db.Where("username = ?", args[3]).First(&user).Error; err != nil {
			fmt.Println("User not found:", args[3])
			return
		}
		task := Task{Name: args[0], Description: args[1], Status: args[2], AssignedTo: user}
		if err := AddTask(db, task); err != nil {
			fmt.Println("Failed to add task:", err)
			return
		}
		fmt.Println("Task added successfully")
	},
}

var viewTaskCmd = &cobra.Command{
	Use:   "view",
	Short: "View a task",
	Long:  `View the details of a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for viewing a task
		id, _ := strconv.Atoi(args[0])
		task, err := ViewTask(db, id)
		if err != nil {
			fmt.Println("Failed to view task:", err)
			return
		}
		fmt.Println(task)
	},
}

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  `Update the details of a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for updating a task
		if len(args) != 5 {
			fmt.Println("Invalid number of arguments. Expected 5 arguments: id, name, description, status, assignedTo.")
			return
		}
		id, _ := strconv.Atoi(args[0])
		var user User
		if err := db.Where("username = ?", args[4]).First(&user).Error; err != nil {
			fmt.Println("User not found:", args[4])
			return
		}
		task := Task{Name: args[1], Description: args[2], Status: args[3], AssignedTo: user}
		if err := UpdateTask(db, id, task); err != nil {
			fmt.Println("Failed to update task:", err)
			return
		}
		fmt.Println("Task updated successfully")
	},
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task from the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for deleting a task
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments. Expected 1 argument: id.")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}
		if err := DeleteTask(db, id); err != nil {
			fmt.Println("Failed to delete task:", err)
			return
		}
		fmt.Println("Task deleted successfully")
	},
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Register, login, and manage user profiles.`,
}

func init() {
	userCmd.AddCommand(registerCmd)
	userCmd.AddCommand(loginCmd)
	userCmd.AddCommand(manageProfileCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Long:  `Register a new user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("Invalid number of arguments. Expected 3 arguments: username, password, profile.")
			return
		}
		username, password, profile := args[0], args[1], args[2]
		user := User{Username: username, Password: password, Profile: profile}
		Register(db, user)
		fmt.Println("User registered successfully")
	},
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login a user",
	Long:  `Login a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for logging in a user
		username := args[0]
		password := args[1]
		if Login(username, password) {
			fmt.Println("Logged in successfully")
		} else {
			fmt.Println("Failed to log in")
		}
	},
}

var manageProfileCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage user profile",
	Long:  `Manage the profile of a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for managing a user profile
		id, _ := strconv.Atoi(args[0])
		username := args[1]
		password := args[2]
		profile := args[3]
		updatedUser := User{Username: username, Password: password, Profile: profile}
		ManageProfile(id, updatedUser)
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with users",
	Long:  `Send and receive chat messages.`,
}

func init() {
	chatCmd.AddCommand(sendMessageCmd)
	chatCmd.AddCommand(receiveMessageCmd)
}

var sendMessageCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message",
	Long:  `Send a chat message to a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for sending a chat message
		from := User{Username: args[0]}
		to := User{Username: args[1]}
		// Since we don't have an actual websocket connection in the CLI context,
		// we'll pass a nil pointer for now. This will need to be handled properly
		// when integrating with a real websocket connection.
		chat := NewChat(0, from, to, nil)
		message := args[2]
		chat.SendMessage(message)
	},
}

var receiveMessageCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive a message",
	Long:  `Receive a chat message from a user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// code for receiving a chat message
		from := User{Username: args[0]}
		to := User{Username: args[1]}
		// Since we don't have an actual websocket connection in the CLI context,
		// we'll pass a nil pointer for now. This will need to be handled properly
		// when integrating with a real websocket connection.
		chat := NewChat(0, from, to, nil)
		stopChan := make(chan struct{}) // Create a channel to signal when to stop receiving messages
		defer close(stopChan)           // Ensure the channel is closed when the function returns
		if err := chat.ReceiveMessage(stopChan); err != nil {
			fmt.Printf("Error receiving messages: %v\n", err)
		}
	},
}
