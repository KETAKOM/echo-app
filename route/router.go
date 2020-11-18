package route

import (
	"net/http"

	"github.com/KETAKOM/echo-app/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", hello)
		v1.GET("/todo", handler.TodoList)
	}

	v2 := e.Group("/api/v2")
	{
		v2.GET("/", hello2)
	}
	return e
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func hello2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World2!")
}
