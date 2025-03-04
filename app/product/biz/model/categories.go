package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Categories struct {
	Base
	Name string `gorm:"column:name"`
}

func (Categories) TableName() string {
	return "category"
}

// GetByCategoryName 根据分类名查询分类
func GetByCategoryName(db *gorm.DB, ctx context.Context, name string) (category *Categories, err error) {
	err = db.WithContext(ctx).Where("name = ?", name).First(&category).Error
	return
}

// GetCategoryNameById 通过标签id查询标签名
func GetCategoryNameById(db *gorm.DB, id []int64) (name []string, err error) {
	var categories []string
	err = db.Table("category").Select("name").Where("id in ?", id).Find(&categories).Error
	if err != nil {
		return name, fmt.Errorf("failed to find categories: %w", err)
	}
	return categories, nil
}

// AddCategories 添加标签
func AddCategories(db *gorm.DB, cat []string) (err error) {
	for _, cat := range cat {
		err = db.Create(&Categories{Name: cat}).Error
		if err != nil {
			fmt.Println(cat, "标签添加失败，", err)
		}
	}
	fmt.Println("标签添加成功")
	return nil
}
