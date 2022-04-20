package routes

import (
	"net/http"
	todo_controller "praktikum-paa-b-2022/controllers/todo"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello world",
		})
	})

	e.GET("/todos", todo_controller.GetAll)
	e.POST("/todos", todo_controller.Create)
	e.PATCH("/todos/:id", todo_controller.Update)
	e.DELETE("/todos/:id", todo_controller.Delete)

	// todoRoute := e.Group("/todos")
	// todoRoute.GET("/", todo_controller.GetAll)
	// todoRoute.POST("/", todo_controller.Create)

	return e
}
