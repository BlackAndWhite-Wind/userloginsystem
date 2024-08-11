package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId      uint   `gorm:"primaryKey"`
	UserName    string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
