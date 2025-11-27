package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique; not null"`
	Password  string    `json:"-" gorm:"not null"`
	Notes     []*Note   `json:"notes" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
}
