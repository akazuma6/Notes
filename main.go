package main

import (
	"go-snsbackend/db"

	"github.com/labstack/echo/v4"
)

func main() {
	// dbパッケージからInitを呼び出し
	db.Init()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
