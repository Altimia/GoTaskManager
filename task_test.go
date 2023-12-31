package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	assert.NoError(t, err)
	defer gormDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `tasks`").WithArgs("Test Task", "Test Description", "Pending", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	task := Task{Name: "Test Task", Description: "Test Description", Status: "Pending", AssignedTo: User{Model: gorm.Model{ID: 1}}}
	err = AddTask(task)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestViewTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	assert.NoError(t, err)
	defer gormDB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "status", "assigned_to_id"}).
		AddRow(1, "Test Task", "Test Description", "Pending", 1)
	mock.ExpectQuery("SELECT * FROM `tasks` WHERE").WithArgs(1).WillReturnRows(rows)

	task, err := ViewTask(gormDB, 1)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Test Task", task.Name)
	assert.Equal(t, "Test Description", task.Description)
	assert.Equal(t, "Pending", task.Status)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	assert.NoError(t, err)
	defer gormDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `tasks` SET").WithArgs("Updated Task", "Updated Description", "Completed", sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	task := Task{Name: "Updated Task", Description: "Updated Description", Status: "Completed", AssignedTo: User{Model: gorm.Model{ID: 1}}}
	err = UpdateTask(1, task)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestDeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open("sqlite3", db)
	assert.NoError(t, err)
	defer gormDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `tasks` WHERE").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	DeleteTask(1)

	assert.NoError(t, mock.ExpectationsWereMet())
}
