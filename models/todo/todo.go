package todo

import (
	"praktikum-paa-b-2022/config"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func CreateTodo(todo *Todo) error {
	err := config.DB.Create(todo).Error

	return err
}

func GetAllTodo() []Todo {
	var todos []Todo

	config.DB.Find(&todos)

	return todos
}

func UpdateTodo(todo *Todo) error {
	err := config.DB.Model(todo).Select("title", "description").Updates(todo).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id int) error {
	err := config.DB.Delete(&Todo{}, id).Error

	return err
}
