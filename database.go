package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// OpenDatabase is an exported function that initializes the database and returns a *gorm.DB instance.
func OpenDatabase() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.AutoMigrate(&Task{}, &User{}, &Chat{})

	return db, nil
}

func InitDB() {
	db, _ = OpenDatabase()

	fmt.Println("Database connected")
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		fmt.Println("Failed to close the database")
		panic("Failed to close the database")
	}
	fmt.Println("Database closed")
}
