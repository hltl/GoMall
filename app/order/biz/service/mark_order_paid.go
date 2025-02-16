package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/order/biz/dal/mysql"
	"github.com/hltl/GoMall/app/order/biz/model"
	order "github.com/hltl/GoMall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		var o model.Order
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("order_id = ?", req.OrderId).First(&o).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return kerrors.NewGRPCBizStatusError(6004001, "no such order")
			}
			return err
		}

		if req.UserId != o.UserID {
			return kerrors.NewGRPCBizStatusError(6004001, "user not match")
		}

		if o.Status == model.OrderStatusPaid {
			return kerrors.NewGRPCBizStatusError(6004001, "order already paid")
		}
		if o.Status == model.OrderStatusCancelled {
			return kerrors.NewGRPCBizStatusError(6004001, "order already cancelled")
		}

		if err := tx.Model(&model.Order{}).Where("order_id = ?", o.OrderID).Update("status", model.OrderStatusPaid).Error; err != nil {
			return kerrors.NewGRPCBizStatusError(6005001, "update order status failed")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &order.MarkOrderPaidResp{}, nil
}
