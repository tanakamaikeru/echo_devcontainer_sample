package controller

import (
	"net/http"

	"myapp/controller/request"

	"myapp/usecase"

	"github.com/labstack/echo/v4"
)

func CreateAuthor(c echo.Context) error {
	body := request.CreateAuthorRequest{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	author := usecase.CreateAuthor(body.Name, body.Age)
	return c.JSON(http.StatusCreated, author)
}

func GetAuthor(c echo.Context) error {
	body := request.GetAuthorRequest{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	err, author := usecase.GetAuthor(body.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found.")
	}
	return c.JSON(http.StatusOK, author)
}
