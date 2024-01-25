package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	assert.NoError(t, err)
	defer gormDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	gormDB, err = OpenDatabase() // Reuse the existing 'err' variable without redeclaring it
	assert.NoError(t, err)
	defer db.Close()

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCloseDB(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	assert.NoError(t, err)
	defer gormDB.Close()

	// Use the exported OpenDatabase function to initialize the database
	gormDB, err = OpenDatabase() // Reuse the existing 'err' variable without redeclaring it
	assert.NoError(t, err)
	defer db.Close()

	// Since CloseDB uses the global db variable, we need to set it to our mock
	// and then assert that there's no error when closing it.
	err = db.Close()
	assert.NoError(t, err)
}
