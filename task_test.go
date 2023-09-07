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

	viewedTask := ViewTask(1)
	if err != nil {
		t.Errorf("Error viewing task: %v", err)
	}
	if viewedTask.Name != "Test Task" {
		t.Errorf("Expected task name to be 'Test Task', but got '%v'", viewedTask.Name)
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

	viewedTask := ViewTask(1)
	if err != nil {
		t.Errorf("Error viewing task: %v", err)
	}
	if viewedTask.Name != "Updated Test Task" {
		t.Errorf("Expected task name to be 'Updated Test Task', but got '%v'", viewedTask.Name)
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

	viewedTask := ViewTask(1)
	if err != nil {
		t.Errorf("Error viewing task: %v", err)
	}
	if viewedTask.ID != 0 {
		t.Errorf("Expected task to be deleted, but it still exists with ID %v", viewedTask.ID)
	}
}

// END: 8f7e6d5b3a4c
