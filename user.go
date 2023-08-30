package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Profile  string
}

func Register(user User) {
	db.Create(&user)
	fmt.Println("User registered successfully")
}

func Login(username string, password string) bool {
	var user User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false
	}
	return true
}

func ManageProfile(id int, updatedUser User) {
	var user User
	db.First(&user, id)
	db.Model(&user).Updates(updatedUser)
	fmt.Println("Profile updated successfully")
}
