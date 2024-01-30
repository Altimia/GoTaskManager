package main

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Status      string
	AssignedTo  User
}

func AddTask(gormDB *gorm.DB, task Task) error {
	if err := gormDB.Create(&task).Error; err != nil {
		zap.L().Error("Error adding task", zap.Error(err))
		return err
	}
	zap.L().Info("Task added successfully")
	return nil
}

func ViewTask(gormDB *gorm.DB, id int) (Task, error) {
	var task Task
	if err := gormDB.First(&task, id).Error; err != nil {
		zap.L().Error("Error viewing task", zap.Int("id", id), zap.Error(err))
		return Task{}, err
	}
	zap.L().Info("Task viewed successfully", zap.Int("id", id))
	return task, nil
}

func UpdateTask(id int, updatedTask Task) error {
func UpdateTask(gormDB *gorm.DB, id int, updatedTask Task) error {
	var task Task
	if err := gormDB.First(&task, id).Error; err != nil {
		zap.L().Error("Error finding task for update", zap.Int("id", id), zap.Error(err))
		return err
	}
	if err := gormDB.Model(&task).Updates(updatedTask).Error; err != nil {
		zap.L().Error("Error updating task", zap.Int("id", id), zap.Error(err))
		return err
	}
	// Send notification to the assigned user if they are connected and the connection is not nil
	userConnectionsMutex.Lock()
	conn, ok := userConnections[task.AssignedTo.ID]
	userConnectionsMutex.Unlock()
	if ok && conn != nil {
		notification := fmt.Sprintf("Task '%s' has been updated.", task.Name)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(notification)); err != nil {
			zap.L().Error("Error sending update notification", zap.Error(err))
		}
	}
	zap.L().Info("Task updated successfully", zap.Int("id", id))
	return nil
}

func DeleteTask(id int) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		zap.L().Error("Error finding task for deletion", zap.Int("id", id), zap.Error(err))
		return
	}
	if err := db.Delete(&task).Error; err != nil {
		zap.L().Error("Error deleting task", zap.Int("id", id), zap.Error(err))
		return
	}
	zap.L().Info("Task deleted successfully", zap.Int("id", id))
}
