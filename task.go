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

func ViewTask(id int) (Task, error) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func UpdateTask(id int, updatedTask Task) error {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}
	if err := db.Model(&task).Updates(updatedTask).Error; err != nil {
		return err
	}
	fmt.Println("Task updated successfully")
	return nil
}

func DeleteTask(id int) {
	var task Task
	db.First(&task, id)
	db.Delete(&task)
	fmt.Println("Task deleted successfully")
}
