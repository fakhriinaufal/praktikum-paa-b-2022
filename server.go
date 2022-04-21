package main

import (
	"praktikum-paa-b-2022/config"
	"praktikum-paa-b-2022/middlewares"
	todo_model "praktikum-paa-b-2022/models/todo"
	"praktikum-paa-b-2022/routes"

	"github.com/spf13/viper"
)

func migrate() {
	config.DB.AutoMigrate(todo_model.Todo{})
}

func init() {
	viper.SetConfigFile("config/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func main() {
	config.InitDB()

	migrate()

	confJWT := middlewares.ConfigJwt{
		Secret:    viper.GetString("jwt.secret"),
		ExpiresAt: viper.GetInt64("jwt.expired"),
	}

	e := routes.New(confJWT.Init())

	e.Start(viper.GetString("server.address"))
}
