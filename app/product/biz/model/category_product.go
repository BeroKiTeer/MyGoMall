package model

import (
	"fmt"
	"gorm.io/gorm"
)

type CategoryProduct struct {
	Id         int64 `gorm:"primary_key;column:id"`
	CategoryId int64 `gorm:"column:category_id"`
	ProductId  int64 `gorm:"column:product_id"`
}

func (CategoryProduct) TableName() string {
	return "category_product"
}

// CreateCPRelation 添加商品和分类
func CreateCPRelation(db *gorm.DB, cp *CategoryProduct) error {
	return db.Create(cp).Error
}

// SelectCategoryId 查询产品类别ID
func SelectCategoryId(db *gorm.DB, id int64) ([]int64, error) {
	var categoriesProduct []int64
	err := db.Table("category_product").Select("category_id").Where("product_id= ?", id).Find(&categoriesProduct).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find product categories: %w", err)
	}
	return categoriesProduct, nil
}

// SelectProductIdByCategoryId 根据标签id查询商品id
func SelectProductIdByCategoryId(db *gorm.DB, page int, pageSize int, categoriesId int) ([]int64, error) {
	var productsId []int64
	err := db.Table("category_product").Select("product_id").Where("category_id = ?", categoriesId).Offset(int(page) - 1).Limit(pageSize).Scan(&productsId).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find products id by category id %d: %w", categoriesId, err)
	}
	return productsId, nil
}
