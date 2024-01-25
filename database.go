package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Task{}, &User{}, &Chat{})

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
