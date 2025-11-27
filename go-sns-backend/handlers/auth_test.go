package handlers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 開発中にコードの一部を確認するためのテスト関数
// 実行方法: go test -run TestDebugPassword -v
func ExampleDebugPassword() {
	password := "mypassword123"

	// バイト配列に変換
	passwordBytes := []byte(password)

	// 内容を確認
	fmt.Printf("元のパスワード (文字列): %s\n", password)
	fmt.Printf("元のパスワード (バイト配列): %v\n", passwordBytes)
	fmt.Printf("バイト配列の長さ: %d\n", len(passwordBytes))

	// ハッシュ化
	hashed, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	fmt.Printf("ハッシュ化されたパスワード: %s\n", string(hashed))
	fmt.Printf("ハッシュの長さ: %d\n", len(hashed))

	// 検証
	err = bcrypt.CompareHashAndPassword(hashed, passwordBytes)
	if err != nil {
		fmt.Printf("検証失敗: %v\n", err)
	} else {
		fmt.Println("検証成功: パスワードが一致しました")
	}
}
