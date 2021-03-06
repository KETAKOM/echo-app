package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/KETAKOM/echo-app/handler/todo"
)

func Init() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/todo", todo.TodoList)
		v1.POST("/todo", todo.CreateTodo)
		v1.PUT("/todo", todo.UpdateTodo)
		v1.DELETE("/todo", todo.DeleteTodo)
	}
	return e
}
