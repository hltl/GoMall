package model

import "gorm.io/gorm"

type Consignee struct {
	Email   string
	Street  string
	City    string
	State   string
	Country string
	ZipCode int32
}

type Order struct {
	gorm.Model
	OrderID      string      `json:"orderId" gorm:"type:varchar(100);unique_index"`
	UserID       uint        `json:"userId"`
	UserCurrency string      `json:"userCurrency"`
	Email        string      `json:"email"`
	Consignee    Consignee   `json:"consignee" gorm:"embedded"`
	Items        []OrderItem `json:"items"`
}

func (Order) TableName() string {
	return "order"
}
