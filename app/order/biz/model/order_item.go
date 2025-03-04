package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductID uint32  `json:"productId"`
	OrderID   string  `json:"orderId" gorm:"size:256;index"`
	Quantity  int32  `json:"quantity"`
	Cost      float32 `json:"cost" gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
