package model

import (
	"time"
)

type Product struct {
	Id            int       `gorm:"primary_key;column:id"`
	CategoryId    int       `gorm:"column:category_id;"`
	Name          string    `gorm:"column:name"`
	Description   string    `gorm:"column:description"`
	Price         float32   `gorm:"column:price"`
	OriginalPrice float32   `gorm:"column:original_price"`
	Images        string    `gorm:"column:images"`
	SalesCount    int       `gorm:"column:sales_count"`
	Status        int       `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at"`
}

func (Product) TableName() string {
	return "products"
}
