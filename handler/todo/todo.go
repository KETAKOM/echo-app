package todo

import (
	"net/http"

	"github.com/KETAKOM/echo-app/model"
	"github.com/labstack/echo"
)

func TodoList(c echo.Context) error {
	todoList, err := model.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoList)
}

func CreateTodo(c echo.Context) error {
	t := new(model.Todo)
	if err := c.Bind(t); err != nil {
		return err
	}

	err := model.CreateTodo(t)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "OK")
}
