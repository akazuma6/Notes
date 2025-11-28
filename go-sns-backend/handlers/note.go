package handlers

import (
	"go-snsbackend/db"
	"go-snsbackend/models"
	"net/http"

	// "github.com/golang-jwt/jwt/v5" // 認証不要のためコメントアウト
	"github.com/labstack/echo/v4"
)

func CreateNote(c echo.Context) error {
	// JWT認証不要: user_idはnull許容
	// jwtミドルウェアからuseridを取得（コメントアウト）
	// c.Get("user")でミドルウェアが保存したJWTトークンを取得
	// userToken := c.Get("user").(*jwt.Token) // !!!型アサーション!!!
	// トークンからClaims（ペイロード）を取得
	// claims := userToken.Claims.(jwt.MapClaims)
	// Claimsからuser_idを取得してuint型に変換
	// userID := claims["user_id"].(uint)
	note := new(models.Note)
	if err := c.Bind(note); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	// note.UserID = userID // ノートにtokenに含まれたユーザーIDをセット（コメントアウト）
	// user_idはリクエストボディから取得されるか、nullのまま
	if err := db.DB.Create(note).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, note) // 成功したら作成したノートを返却
}

func GetNotes(c echo.Context) error {
	var notes []models.Note

	if err := db.DB.Preload("User").Order("created_at desc").Find(&notes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.JSON(http.StatusOK, notes)
}
