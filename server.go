package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"myapp/controller"
)

func main() {
	e := echo.New()
	e.POST("/authors", controller.CreateAuthor)
	e.GET("/authors/:id", controller.GetAuthor)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
