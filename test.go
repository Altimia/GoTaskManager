package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// Mock database and expectations
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Set up the mock expectations
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WithArgs("testuser", "testpass", "testprofile").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Use the mocked DB connection
	gormDB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}
	defer func() { gormDB.Close() }()

	// Test the Register function
	Register(User{Username: "testuser", Password: "testpass", Profile: "testprofile"})

	// Ensure all expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLogin(t *testing.T) {
	// Mock database and expectations
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Set up the mock expectations
	rows := sqlmock.NewRows([]string{"id", "username", "password", "profile"}).
		AddRow(1, "testuser", "testpass", "testprofile")
	mock.ExpectQuery("SELECT * FROM `users`").WithArgs("testuser", "testpass").WillReturnRows(rows)

	// Use the mocked DB connection
	gormDB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}
	defer func() { gormDB.Close() }()

	// Test the Login function
	result := Login("testuser", "testpass")

	// Check the result
	assert.True(t, result)

	// Ensure all expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Additional tests for task management and real-time communication would follow a similar pattern
