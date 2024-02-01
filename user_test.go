package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestRegister(t *testing.T) {
	// Mock database and expectations
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	// Set up the mock expectations
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"users\" \\(\"created_at\",\"updated_at\",\"deleted_at\",\"username\",\"password\",\"profile\"\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?\\)").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "testuser", "testpass", "testprofile").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Use the mocked DB connection
	gormDB, err := gorm.Open("sqlite3", sqlDB)
	assert.NoError(t, err)
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}
	defer func() { gormDB.Close() }()

	// Inject the mocked gormDB into the Register function
	err = Register(gormDB, User{Username: "testuser", Password: "testpass", Profile: "testprofile"})
	assert.NoError(t, err)

	// Test the Register function
	Register(gormDB, User{Username: "testuser", Password: "testpass", Profile: "testprofile"})

	// Ensure all expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLogin(t *testing.T) {
	// Mock database and expectations
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Set up the mock expectations
	expectedSQL := "SELECT \\* FROM \"users\" WHERE \\(username = \\? AND password = \\?\\) ORDER BY \"users\".\"id\" ASC LIMIT 1"
	currentTime := time.Now()
	rows := sqlmock.NewRows([]string{"id", "username", "password", "profile", "created_at", "updated_at"}).
		AddRow(1, "testuser", "testpass", "testprofile", currentTime, currentTime)
	mock.ExpectQuery(expectedSQL).WithArgs("testuser", "testpass").WillReturnRows(rows)

	// Use the mocked DB connection
	assert.NoError(t, err)
	gormDB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}
	defer func() { gormDB.Close() }()

	// Test the Login function with the correct arguments
	result := Login(gormDB, "testuser", "testpass")

	// Check the result
	assert.True(t, result, "Login should return true for valid credentials")

	// Ensure all expectations are met
	assert.NoError(t, mock.ExpectationsWereMet())
}
