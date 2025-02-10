package model

import (
	"time"
)

type Product struct {
	//gorm.Model
	Id            int       `gorm:"primaryKey;type:bigint;"`
	CategoryId    int       `gorm:"type:bigint;"`
	Name          string    `gorm:"type:varchar(100);"`
	Description   string    `gorm:"type:text;"`
	Price         float32   `gorm:"type:decimal(10,2);"`
	OriginalPrice float32   `gorm:"type:decimal(10,2);"`
	Images        string    `gorm:"type:json;"`
	SalesCount    int       `gorm:"type:int;"`
	Status        int       `gorm:"type:tinyint;"`
	CreatedAt     time.Time `gorm:"type:timestamp;"`
	UpdatedAt     time.Time `gorm:"type:timestamp;"`
	DeletedAt     time.Time `gorm:"type:timestamp;"`
}

func (Product) TableName() string {
	return "products"
}
