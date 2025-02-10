package model

import "gorm.io/gorm"

type Product struct {
	Base
	CategoryId    int     `json:"category_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Stock         int     `json:"stock"`
	Images        string  `json:"images"`
	Status        int     `json:"status"`
}

func (p Product) TableName() string {
	return "products"
}

// 添加商品
func AddProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}
