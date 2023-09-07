package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestAddTaskWithInvalidUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("sqlite3", db)

	// Mock the query
	mock.ExpectBegin()
	mock.ExpectQuery("^INSERT INTO \"tasks\" (.+)$").WillReturnError(gorm.ErrRecordNotFound)

	task := Task{
		Name:        "Test Task",
		Description: "This is a test task",
		Status:      "In Progress",
		AssignedTo: User{
			Username: "Invalid User",
			Password: "password123",
		},
	}

	err := AddTask(task)

	// Check that the error is what we expected
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

// Add more tests here...
