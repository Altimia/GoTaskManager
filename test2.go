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

	_, _ = gorm.Open("sqlite3", db)

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

	err := gormDB.Create(&task).Error

	// Check that the error is what we expected
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

// Add more tests here...
func TestViewTask2(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("sqlite3", db)

	// Mock the query
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM \"tasks\" WHERE \"tasks\".\"id\" = ? ORDER BY \"tasks\".\"id\" ASC LIMIT 1$").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "status", "assigned_to"}).AddRow(1, "Test Task", "This is a test task", "In Progress", "John Doe"))

	task := ViewTask(1)

	// Check that the task is what we expected
	assert.Equal(t, "Test Task", task.Name)
	assert.Equal(t, "This is a test task", task.Description)
	assert.Equal(t, "In Progress", task.Status)
	assert.Equal(t, "John Doe", task.AssignedTo.Username)
}

func TestUpdateTask2(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("sqlite3", db)

	// Mock the query
	mock.ExpectBegin()
	mock.ExpectQuery("^UPDATE \"tasks\" SET \"name\" = ?, \"description\" = ?, \"status\" = ?, \"assigned_to\" = ? WHERE \"tasks\".\"id\" = ?$").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "status", "assigned_to"}).AddRow(1, "Updated Test Task", "This is an updated test task", "Completed", "Jane Doe"))

	updatedTask := Task{
		Name:        "Updated Test Task",
		Description: "This is an updated test task",
		Status:      "Completed",
		AssignedTo: User{
			Username: "Jane Doe",
			Password: "password123",
		},
	}

	UpdateTask(1, updatedTask)

	task := ViewTask(1)

	// Check that the task was updated correctly
	assert.Equal(t, "Updated Test Task", task.Name)
	assert.Equal(t, "This is an updated test task", task.Description)
	assert.Equal(t, "Completed", task.Status)
	assert.Equal(t, "Jane Doe", task.AssignedTo.Username)
}

func TestDeleteTask2(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open("sqlite3", db)

	// Mock the query
	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM \"tasks\" WHERE \"tasks\".\"id\" = ?$").WillReturnResult(sqlmock.NewResult(1, 1))

	DeleteTask(1)

	task := ViewTask(1)

	// Check that the task was deleted correctly
	assert.Equal(t, 0, task.ID)
}
