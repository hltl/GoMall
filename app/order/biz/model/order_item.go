package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductID uint32  `json:"productId"`
	OrderID   string  `json:"orderId" gorm:"type:varchar(100)"`
	Quantity  uint32  `json:"quantity"`
	Cost      float64 `json:"cost" gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
