package db

// アプリ起動時に自動でテーブルを作る

import (
	"fmt"
	"go-snsbackend/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// ユーザー:パスワード@tcp(ホスト:ポート)/DB名
	// 接続情報の作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?cahrset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error
	// データベースに接続
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
	// 自動的にMySQL上にテーブルを作成
	DB.AutoMigrate(&models.User{}, &models.Note{})
}
