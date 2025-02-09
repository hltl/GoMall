package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

func GetProductById(ctx context.Context, db *gorm.DB, id int) (product Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Where(&Product{Base: Base{ID: id}}).First(&product).Error
	return
}

func SearchProduct(ctx context.Context, db *gorm.DB, q string) (products []*Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}
