package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionID string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func Pay(ctx context.Context, db *gorm.DB,pl *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(pl).Error
}
