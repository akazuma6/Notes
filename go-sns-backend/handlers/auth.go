package handlers

import (
	"go-snsbackend/db"
	"go-snsbackend/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to bind user"})
	}
	// []byte(user.Password)はパスワードをバイト配列に変換
	// デバッグ: パスワードの内容を確認（開発中のみ）
	// fmt.Printf("DEBUG: 元のパスワード (文字列): %s\n", user.Password)
	// fmt.Printf("DEBUG: 元のパスワード (バイト配列): %v\n", []byte(user.Password))
	// fmt.Printf("DEBUG: バイト配列の長さ: %d\n", len([]byte(user.Password)))

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// デバッグ: ハッシュ化後の内容を確認
	// fmt.Printf("DEBUG: ハッシュ化されたパスワード: %s\n", string(hashed))
	// fmt.Printf("DEBUG: ハッシュの長さ: %d\n", len(hashed))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not hash password")
	}
	// ハッシュ化したパスワードを再セット
	user.Password = string(hashed)

	// データベースに保存
	if err := db.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}
	// パスワードを除外してレスポンス
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to bind login request"})
	}
	var user models.User
	// 今作ったuserに書き込むために&をつける
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	// パスワード照合
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Could not generate token")
	}

	//JWTトークンを生成
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
