package todo

import (
	"net/http"
	"praktikum-paa-b-2022/controllers"
	"praktikum-paa-b-2022/middlewares"
	"praktikum-paa-b-2022/models/todo"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GetAll(c echo.Context) error {
	todos := todo.GetAllTodo()

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"todos": todos,
	// })
	return controllers.NewSuccessResponse(c, todos)
}

func Create(c echo.Context) error {
	var newTodo todo.Todo

	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error happened",
		})
	}

	err := todo.CreateTodo(&newTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error happened",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"todo": newTodo,
	})
}

func Update(c echo.Context) error {
	var updateTodo todo.Todo

	if err := c.Bind(&updateTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error happened",
		})
	}

	id := c.Param("id")
	convertedId, _ := strconv.Atoi(id)

	updateTodo.ID = convertedId

	err := todo.UpdateTodo(&updateTodo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error happened",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"todo": updateTodo,
	})
}

func Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := todo.DeleteTodo(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error happened",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "todo deleted",
	})
}

func GenerateTokenHTTP(c echo.Context) error {
	confJWT := middlewares.ConfigJwt{
		Secret:    viper.GetString("jwt.secret"),
		ExpiresAt: viper.GetInt64("jwt.expired"),
	}
	tokenn, err := confJWT.GenerateToken("ini title")
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, tokenn)
}
