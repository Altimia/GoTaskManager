package main

import (
	"fmt"
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
		return err
	}
	fmt.Println("Task added successfully")
	return nil
}

func ViewTask(id int) Task {
	var task Task
	db.First(&task, id)
	return task
}

func UpdateTask(id int, updatedTask Task) {
	var task Task
	db.First(&task, id)
	db.Model(&task).Updates(updatedTask)
	fmt.Println("Task updated successfully")
}

func DeleteTask(id int) {
	var task Task
	db.First(&task, id)
	db.Delete(&task)
	fmt.Println("Task deleted successfully")
}
