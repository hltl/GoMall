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

func (pl *PaymentLog) Save(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(pl).Error
}
