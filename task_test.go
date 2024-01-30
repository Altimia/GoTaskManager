package main

import (
	"regexp"
	"testing"

	"time"

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
	mock.ExpectExec("INSERT INTO \"tasks\" \\(\"created_at\",\"updated_at\",\"deleted_at\",\"name\",\"description\",\"status\"\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?\\)").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "Test Task", "Test Description", "Pending").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	task := Task{Name: "Test Task", Description: "Test Description", Status: "Pending", AssignedTo: User{Model: gorm.Model{ID: 1}}}
	err = AddTask(gormDB, task)
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

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "status", "assigned_to_id"}).
		AddRow(1, time.Now(), time.Now(), nil, "Test Task", "Test Description", "Pending", 1)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."deleted_at" IS NULL AND (("tasks"."id" = 1)) ORDER BY "tasks"."id" ASC LIMIT 1`)).WillReturnRows(rows)

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

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."deleted_at" IS NULL AND (("tasks"."id" = 1)) ORDER BY "tasks"."id" ASC LIMIT 1`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "status", "assigned_to_id"}).AddRow(1, time.Now(), time.Now(), nil, "Existing Task", "Existing Description", "Pending", 1))
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tasks" SET "description" = ?, "name" = ?, "status" = ?, "updated_at" = ? WHERE "tasks"."deleted_at" IS NULL AND "tasks"."id" = ?`)).WithArgs("Updated Description", "Updated Task", "Completed", sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	task := Task{Name: "Updated Task", Description: "Updated Description", Status: "Completed", AssignedTo: User{Model: gorm.Model{ID: 1}}}
	err = UpdateTask(gormDB, 1, task)
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

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."deleted_at" IS NULL AND (("tasks"."id" = 1)) ORDER BY "tasks"."id" ASC LIMIT 1`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "status", "assigned_to_id"}).AddRow(1, time.Now(), time.Now(), nil, "Test Task", "Test Description", "Pending", 1))
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tasks" SET "deleted_at"=? WHERE "tasks"."deleted_at" IS NULL AND "tasks"."id" = ?`)).WithArgs(sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = DeleteTask(gormDB, 1)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
