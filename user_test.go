package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
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
