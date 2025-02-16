package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email   string
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model
	OrderID      string      `json:"orderId" gorm:"type:varchar(100);unique_index"`
	UserID       uint32      `json:"userId"`
	UserCurrency string      `json:"userCurrency"`
	Email        string      `json:"email"`
	Consignee    Consignee   `json:"consignee" gorm:"embedded"`
	Items        []OrderItem `json:"items"`
	Status       OrderStatus `json:"status"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	if err := db.WithContext(ctx).
		Where("user_id = ?", userId).
		Preload("Items").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrder(ctx context.Context, db *gorm.DB, orderId string) (*Order, error) {
    var order Order
    if err := db.WithContext(ctx).
        Where("order_id = ?", orderId).
        First(&order).Error; err != nil {
        return nil, err
    }
    return &order, nil
}
