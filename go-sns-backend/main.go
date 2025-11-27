package main

import (
	"go-snsbackend/db"
	"go-snsbackend/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// dbパッケージからInitを呼び出し
	db.Init()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// 認証関連のエンドポイント
	e.POST("/signup", handlers.Signup)
	e.POST("/login", handlers.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
