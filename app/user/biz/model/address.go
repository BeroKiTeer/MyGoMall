package model

import (
	"context"
	"gorm.io/gorm"
)

type Address struct {
	Id          int64
	UserId      int32 `gorm:"unique"`
	AreaId      int64
	Address     string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	isDeleted   bool
	Status      int8
	Role        int8
}

func (a Address) TableName() string {
	return "address"
}

func GetAddressById(db *gorm.DB, ctx context.Context, id int64) (address *Address, err error) {
	err = db.WithContext(ctx).Where("id = ?", id).First(&address).Error
	return address, err
}

func GetAddressByUserId(db *gorm.DB, ctx context.Context, userId int32) (address *Address, err error) {
	err = db.WithContext(ctx).Where("user_id = ?", userId).First(&address).Error
	return address, err
}
