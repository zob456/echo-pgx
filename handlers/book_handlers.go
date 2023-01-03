package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zob456/echo-pgx/data"
	"net/http"
)

func GetBook(c echo.Context) error {
	id := c.Param("id")
	book, err := data.SelectBook(id, context.Background())
	if err != nil {
		return err
	}
	err = c.JSON(http.StatusOK, book)
	if err != nil {
		return err
	}
	return nil
}
