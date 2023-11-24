package main

import (
	"fmt"
	"log"

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

func AddTask(task Task) error {
	if err := db.Create(&task).Error; err != nil {
		log.Printf("Error adding task: %v", err)
		return err
	}
	log.Println("Task added successfully")
	return nil
}

func ViewTask(id int) (Task, error) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		log.Printf("Error viewing task with id %d: %v", id, err)
		return Task{}, err
	}
	log.Printf("Task with id %d viewed successfully", id)
	return task, nil
}

func UpdateTask(id int, updatedTask Task) error {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		log.Printf("Error finding task with id %d for update: %v", id, err)
		return err
	}
	if err := db.Model(&task).Updates(updatedTask).Error; err != nil {
		log.Printf("Error updating task with id %d: %v", id, err)
		return err
	}
	// Send notification to the assigned user if they are connected
	if conn, ok := userConnections[task.AssignedTo.ID]; ok {
		notification := fmt.Sprintf("Task '%s' has been updated.", task.Name)
		conn.WriteMessage(websocket.TextMessage, []byte(notification))
	}
	fmt.Println("Task updated successfully")
	return nil
}

func DeleteTask(id int) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		log.Printf("Error finding task with id %d for deletion: %v", id, err)
		return
	}
	if err := db.Delete(&task).Error; err != nil {
		log.Printf("Error deleting task with id %d: %v", id, err)
		return
	}
	log.Printf("Task with id %d deleted successfully", id)
}
