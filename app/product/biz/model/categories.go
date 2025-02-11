package model

type Categories struct {
	Base
	Name string `gorm:"column:name"`
}

func (Categories) TableName() string {
	return "categories"
}
