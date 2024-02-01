package main

import (
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Profile  string
}

func Register(gormDB *gorm.DB, user User) error {
	if err := gormDB.Create(&user).Error; err != nil {
		zap.L().Error("Error registering user", zap.Error(err))
		return err
	}
	zap.L().Info("User registered successfully")
	return nil
}

func Login(gormDB *gorm.DB, username string, password string) bool {
	var user User
	if err := gormDB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		zap.L().Error("Login failed for user", zap.String("username", username), zap.Error(err))
		return false
	}
	zap.L().Info("User logged in successfully", zap.String("username", username))
	return true
}

func ManageProfile(id int, updatedUser User) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		zap.L().Error("Error finding user for profile update", zap.Int("id", id), zap.Error(err))
		return
	}
	if err := db.Model(&user).Updates(updatedUser).Error; err != nil {
		zap.L().Error("Error updating profile for user", zap.Int("id", id), zap.Error(err))
		return
	}
	zap.L().Info("Profile for user updated successfully", zap.Int("id", id))
}
