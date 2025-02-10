package model

type Categories struct {
	Base
	name string `gorm:"column:name"`
}

func (Categories) TableName() string {
	return "categories"
}
