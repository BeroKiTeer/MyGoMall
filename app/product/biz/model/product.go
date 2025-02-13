package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name          string  `gorm:"column:name"`
	Description   string  `gorm:"column:description"`
	Price         float32 `gorm:"column:price"`
	OriginalPrice float32 `gorm:"column:original_price"`
	Images        string  `gorm:"column:images"`
	Stock         uint32  `gorm:"column:stock"`
	Status        int     `gorm:"column:status"`
}

func (p Product) TableName() string {
	return "products"
}

// CreateProduct 添加商品
func CreateProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

func GetProduct(db *gorm.DB, id int) (Product, error) {
	var row Product
	err := db.Model(&Product{}).Where("id=?", id).Find(&row).Error
	if err != nil {
		return row, fmt.Errorf("failed to find product: %w", err)
	}
	return row, nil
}

// GetProductWithCategory 按照id查询单个商品
func GetProductWithCategory(db *gorm.DB, id int) (Product, []string, error) {
	var row Product
	var categories []string

	//检查数据库连接是否有效
	if db == nil {
		return row, categories, fmt.Errorf("database connection is nil")
	}

	// 查询产品信息
	row, err := GetProduct(db, id)

	// 查询产品类别ID
	categoryId, err := SelectCategoryId(db, int64(id))
	if err != nil {
		return Product{}, nil, err
	}

	// 查询类别名称
	categories, err = GetCategoryNameById(db, categoryId)
	if err != nil {
		return Product{}, nil, err
	}
	fmt.Printf("%+v\n%+v\n", row, categories)
	return row, categories, nil
}

// GetProductsByCategoryName 根据标签分页查询商品
func GetProductsByCategoryName(db *gorm.DB, page int, pageSize int, categoryName string) (products []Product, categories [][]string, err error) {
	var categoriesId int
	//根据名称查询标签id
	db.Table("categories").Select("id").Where("name = ?", categoryName).Scan(&categoriesId)
	fmt.Printf("%+v\n", categoriesId)
	var productsId []int64
	//根据标签id查询商品id
	productsId, err = SelectProductIdByCategoryId(db, page, pageSize, categoriesId)
	if err != nil {
		return products, nil, err
	}
	fmt.Printf("%+v\n", productsId)
	for _, item := range productsId {
		p, category, _ := GetProductWithCategory(db, int(item))
		products = append(products, p)
		categories = append(categories, category)
	}
	fmt.Printf("%+v\n", products)
	return products, categories, nil
}

// DeleteProductById 根据id删除商品
func DeleteProductById(db *gorm.DB, id int) error {
	return db.Where("id=?", id).Delete(&Product{}).Error
}

// GetProductByName 根据名称模糊查询商品
func GetProductByName(db *gorm.DB, name string) ([]Product, [][]string, error) {
	var row []Product
	var categories [][]string
	db.Model(&Product{}).Where("name like ?", "%"+name+"%").Find(&row)
	for _, item := range row {
		_, category, _ := GetProductWithCategory(db, item.ID)
		categories = append(categories, category)
	}
	return row, categories, nil
}

// UpdateProduct 更新商品
func UpdateProduct(db *gorm.DB, product *Product) error {
	fmt.Printf("%+v\n", product)
	return db.Exec(`update products 
					   set name=? ,
					       description=?,
					       price=?,
					       original_price=?, 
					       stock=?,
					       images=?,
					       status=? 
					   where id=?`,
		product.Name,
		product.Description,
		product.Price,
		product.OriginalPrice,
		product.Stock,
		product.Images,
		product.Status,
		product.ID).Error
}
