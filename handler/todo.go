package handler

import (
	"net/http"

	"github.com/KETAKOM/echo-app/model"
	"github.com/labstack/echo"
)

func TodoList(c echo.Context) error {
	// db, err := database.Connect()
	// if err != nil {
	// 	return err
	// }
	// defer db.Close()
	// tasks, err := tasks.New(db).FindAll()
	// if err != nil {
	// 	return err
	// }
	// tasksResponse := make([]helper.ResponseMap, len(tasks))
	// for i, task := range tasks {
	// 	tasksResponse[i] = task.ToResponseMap()
	// }
	todoList, err := model.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todoList)
}
