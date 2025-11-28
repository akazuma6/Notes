package main

import (
	"go-snsbackend/db"
	"go-snsbackend/handlers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// dbパッケージからInitを呼び出し
	db.Init()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 認証関連のエンドポイント
	e.POST("/signup", handlers.Signup)
	e.POST("/login", handlers.Login)

	// ノート関連のエンドポイント
	r := e.Group("/notes")

	// jwtmiddlewareの設定
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	// /notesのルートはようログインにする
	r.Use(echojwt.WithConfig(config))

	r.POST("", handlers.CreateNote)
	r.GET("", handlers.GetNotes)

	e.Logger.Fatal(e.Start(":8080"))
}
