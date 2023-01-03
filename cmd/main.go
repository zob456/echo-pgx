package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zob456/echo-pgx/handlers"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e := echo.New()
	e.GET("/:id", handlers.GetBook)
	e.Logger.Fatal(e.Start(":8080"))
}