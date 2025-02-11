package model

import (
	"context"

	"gorm.io/gorm"
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

func AddItem(ctx context.Context, db *gorm.DB, userId uint, item Item) (err error) {
	var c Cart
	if err := db.WithContext(ctx).Model(&Cart{}).
		Where(Cart{UserID: userId,ProductID: item.ProductId}).
		First(&c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			newCart := Cart{
				UserID:    userId,
				ProductID: item.ProductId,
				Quantity:  item.Quantity,
			}
			return db.WithContext(ctx).Create(&newCart).Error
		}
		return err
	}

	c.Quantity += item.Quantity
	c.Quantity = max(c.Quantity, 0)
	return db.WithContext(ctx).Save(&c).Error
}
func EmptyCart(ctx context.Context, db *gorm.DB,userId uint)(err error){
	return db.WithContext(ctx).Where("user_id = ?", userId).Delete(&Cart{}).Error
}

func GetItemsByUserId(ctx context.Context, db *gorm.DB, userId uint)(items []*Item, err error){
	var Carts []Cart
	if err = db.WithContext(ctx).Where("user_id = ?", userId).Find(&Carts).Error; err != nil {
		return nil, err
	}

	items = make([]*Item, len(Carts))
	for i, ci := range Carts {
		items[i] = &Item{
			ProductId: ci.ProductID,
			Quantity:  ci.Quantity,
		}
	}
	return items, nil
}




