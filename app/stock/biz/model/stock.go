package model

import (
	"gorm.io/gorm"
	"stock/biz/dal/mysql"
)

// CheckQuantity 查询商品的数量
func CheckQuantity(ProductID int64) (int64, error) {
	var quantity int64
	err := mysql.DB.Table("products").
		Select("stock").
		Where("id = ?", ProductID).
		Scan(&quantity).Error
	return quantity, err
}

// ReduceItem 减少库存里的商品
func ReduceItem(ProductID int64, Quantity int64) error {
	err := mysql.DB.Table("products").
		Select("stock").
		Where("id = ?", ProductID).
		Update("stock", gorm.Expr("stock - ?", Quantity)).Error
	return err
}
