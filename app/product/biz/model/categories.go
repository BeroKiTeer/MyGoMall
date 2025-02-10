package model

type Categories struct {
	id   int    `gorm:"primary_key;column:id"`
	name string `gorm:"column:name"`
}

func (Categories) TableName() string {
	return "categories"
}
