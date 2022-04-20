package main

import (
	"praktikum-paa-b-2022/config"
	todo_model "praktikum-paa-b-2022/models/todo"
	"praktikum-paa-b-2022/routes"
)

func migrate() {
	config.DB.AutoMigrate(todo_model.Todo{})
}

func main() {
	config.InitDB()

	migrate()

	e := routes.New()

	e.Start(":8080")
}
