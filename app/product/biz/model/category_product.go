package model

type CategoryProduct struct {
	Id         int64 `gorm:"primary_key;column:id"`
	CategoryId int64 `gorm:"column:category_id"`
	ProductId  int64 `gorm:"column:product_id"`
}

func (CategoryProduct) TableName() string {
	return "category_product"
}
