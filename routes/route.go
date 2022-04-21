package routes

import (
	"net/http"
	todo_controller "praktikum-paa-b-2022/controllers/todo"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(authInit middleware.JWTConfig) *echo.Echo {
	e := echo.New()
	jwtauth := middleware.JWTWithConfig(authInit)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello world",
		})
	})

	e.GET("/todos", todo_controller.GetAll)
	e.POST("/todos", todo_controller.Create, jwtauth)
	e.PATCH("/todos/:id", todo_controller.Update)
	e.DELETE("/todos/:id", todo_controller.Delete)
	e.GET("/generateToken", todo_controller.GenerateTokenHTTP)

	// todoRoute := e.Group("/todos")
	// todoRoute.GET("/", todo_controller.GetAll)
	// todoRoute.POST("/", todo_controller.Create)

	return e
}
