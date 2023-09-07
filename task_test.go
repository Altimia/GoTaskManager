package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB()
	code := m.Run()
	CloseDB()
	os.Exit(code)
}

func TestAddTask(t *testing.T) {
	task := Task{
		Name:        "Test Task",
		Description: "This is a test task",
		Status:      "In Progress",
		AssignedTo: User{
			Username: "John Doe",
			Password: "password123",
		},
	}

	err := AddTask(task)
	if err != nil {
		t.Errorf("Error adding task: %v", err)
	}
}

func TestViewTask(t *testing.T) {
	task := Task{
		Name:        "Test Task",
		Description: "This is a test task",
		Status:      "In Progress",
		AssignedTo: User{
			Username: "John Doe",
			Password: "password123",
		},
	}

	err := AddTask(task)
	if err != nil {
		t.Errorf("Error adding task: %v", err)
	}

	_, err = ViewTask(1)
	if err == nil {
		t.Errorf("Expected error when viewing non-existent task, but got no error")
	}
}

func TestUpdateTask(t *testing.T) {
	task := Task{
		Name:        "Test Task",
		Description: "This is a test task",
		Status:      "In Progress",
		AssignedTo: User{
			Username: "John Doe",
			Password: "password123",
		},
	}

	err := AddTask(task)
	if err != nil {
		t.Errorf("Error adding task: %v", err)
	}

	updatedTask := Task{
		Name:        "Updated Test Task",
		Description: "This is an updated test task",
		Status:      "Completed",
		AssignedTo: User{
			Username: "Jane Doe",
			Password: "wordpass13",
		},
	}

	UpdateTask(1, updatedTask)

	_, err = ViewTask(1)
	if err == nil {
		t.Errorf("Expected error when viewing non-existent task, but got no error")
	}
}

func TestDeleteTask(t *testing.T) {
	task := Task{
		Name:        "Test Task",
		Description: "This is a test task",
		Status:      "In Progress",
		AssignedTo: User{
			Username: "John Doe",
			Password: "password123",
		},
	}

	err := AddTask(task)
	if err != nil {
		t.Errorf("Error adding task: %v", err)
	}

	DeleteTask(1)

	_, err = ViewTask(1)
	if err == nil {
		t.Errorf("Expected error when viewing deleted task, but got no error")
	}
}

// END: 8f7e6d5b3a4c
