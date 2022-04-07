package todo

import (
	"net/http"
	"praktikum-paa-b-2022/models/todo"

	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"todos": todo.Todos,
	})
}

func Create(c echo.Context) error {
	var newTodo todo.Todo

	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error happened",
		})
	}

	todo.Todos = append(todo.Todos, newTodo)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"todo": newTodo,
	})
}
