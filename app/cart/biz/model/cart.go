package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cart struct {
	UserID    uint `json:"userId" gorm:"primaryKey"`
	ProductID uint `json:"productId" gorm:"primaryKey"`
	Quantity  int  `json:"quantity"`
}

func (c Cart) TableName() string {
	return "cart"
}

type Item struct {
	ProductId uint
	Quantity  int
}

func AddItem(ctx context.Context, db *gorm.DB, userId uint, item Item) error {
	return db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_id"},
			{Name: "product_id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"quantity": gorm.Expr("GREATEST(quantity + ?, 0)", item.Quantity),
		}),
	}).Create(&Cart{
		UserID:    userId,
		ProductID: item.ProductId,
		Quantity:  item.Quantity,
	}).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint) (err error) {
	return db.WithContext(ctx).Where("user_id = ?", userId).Delete(&Cart{}).Error
}

func GetItemsByUserId(ctx context.Context, db *gorm.DB, userId uint) ([]*Item, error) {
	var tmp []Item
	if err := db.WithContext(ctx).
		Model(&Cart{}).
		Where("user_id = ?", userId).
		Select("product_id, quantity").
		Scan(&tmp).Error; err != nil {
		return nil, err
	}

	items := make([]*Item, len(tmp))
	for i := range tmp {
		items[i] = &tmp[i]
	}
	return items, nil
}