package db

// アプリ起動時に自動でテーブルを作る

import (
	"fmt"
	"go-snsbackend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}

	// ユーザー:パスワード@tcp(ホスト:ポート)/DB名
	// 接続情報の作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// データベースに接続
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to database")
	// 自動的にMySQL上にテーブルを作成
	DB.AutoMigrate(&models.User{}, &models.Note{})

	// 初期データを投入
	Seed()
}
