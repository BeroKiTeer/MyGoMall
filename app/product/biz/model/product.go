package model

import "gorm.io/gorm"

type Product struct {
	Base
	CategoryId    int     `gorm:"column:category_id;"`
	Name          string  `gorm:"column:name"`
	Description   string  `gorm:"column:description"`
	Price         float32 `gorm:"column:price"`
	OriginalPrice float32 `gorm:"column:original_price"`
	Images        string  `gorm:"column:images"`
	Status        int     `gorm:"column:status"`
}

func (p Product) TableName() string {
	return "products"
}

// 添加商品
func AddProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}
