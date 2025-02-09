package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username       string `gorm:"unique"`
	PasswordHashed string
	PhoneNumber    string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Address        string
	Status         bool
	Role           bool
}

func (u User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return
}
