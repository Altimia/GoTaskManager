package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Profile  string
}

func Register(user User) {
	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error registering user: %v", err)
		return
	}
	log.Println("User registered successfully")
}

func Login(username string, password string) bool {
	var user User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		log.Printf("Login failed for user %s: %v", username, err)
		return false
	}
	log.Printf("User %s logged in successfully", username)
	return true
}

func ManageProfile(id int, updatedUser User) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		log.Printf("Error finding user with id %d for profile update: %v", id, err)
		return
	}
	if err := db.Model(&user).Updates(updatedUser).Error; err != nil {
		log.Printf("Error updating profile for user with id %d: %v", id, err)
		return
	}
	log.Printf("Profile for user with id %d updated successfully", id)
}
