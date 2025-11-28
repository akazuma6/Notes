package db

import (
	"go-snsbackend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Seed 初期データを投入する関数
func Seed() {
	// ユーザーの初期データ
	var count int64
	DB.Model(&models.User{}).Count(&count)

	if count == 0 {
		log.Println("初期データを投入します...")

		// 初期ユーザー1: admin
		hashedPassword1, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			return
		}
		adminUser := models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: string(hashedPassword1),
		}
		if err := DB.Create(&adminUser).Error; err != nil {
			log.Printf("初期ユーザーの作成に失敗しました: %v", err)
		} else {
			log.Printf("初期ユーザー '%s' を作成しました", adminUser.Username)

			// 初期ノートを追加
			note1 := models.Note{
				UserID:  adminUser.ID,
				Content: "これは初期ノートです。",
			}
			if err := DB.Create(&note1).Error; err != nil {
				log.Printf("初期ノートの作成に失敗しました: %v", err)
			} else {
				log.Printf("初期ノートを作成しました")
			}
		}
	}
}
