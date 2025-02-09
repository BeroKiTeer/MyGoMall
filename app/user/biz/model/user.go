package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username       string `gorm:"unique"`
	PasswordHashed string
	PhoneNumber    string `gorm:"unique"`
	Email          string `gorm:"unique"`
	Address        string
	Status         int8
	Role           int8
}

func (u User) TableName() string {
	return "user"
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetUserById(db *gorm.DB, ctx context.Context, id int32) (user *User, err error) {
	err = db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return
}

// UserExistsByID 检查用户是否存在
func UserExistsByID(db *gorm.DB, userID int32) (bool, error) {
	var exists bool
	err := db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", userID).
		Scan(&exists).
		Error

	if err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateUser 更新用户信息
func UpdateUser(db *gorm.DB, ctx context.Context, userID int32, updates map[string]interface{}) error {
	// 1️⃣ 检查用户是否存在
	var user User
	err := db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}

	// 2️⃣ 执行更新操作
	err = db.WithContext(ctx).Model(&user).Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}
