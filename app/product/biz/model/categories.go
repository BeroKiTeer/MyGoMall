package model

import (
	"context"
	"gorm.io/gorm"
)

type Categories struct {
	Base
	Name string `gorm:"column:name"`
}

func (Categories) TableName() string {
	return "categories"
}

// 根据分类名查询分类
func GetByCategoryName(db *gorm.DB, ctx context.Context, name string) (category *Categories, err error) {
	err = db.WithContext(ctx).Where("name = ?", name).First(&category).Error
	return
}
