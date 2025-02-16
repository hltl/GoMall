package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"github.com/hltl/GoMall/app/order/biz/dal/mysql"
	"github.com/hltl/GoMall/app/order/biz/model"
	order "github.com/hltl/GoMall/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		return nil, kerrors.NewGRPCBizStatusError(6004001, "order items is required")
	}
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// Create order
		// Create order items
		orderId, _ := uuid.NewUUID()
		o := &model.Order{
			OrderID:      orderId.String(),
			UserID:       req.UserId,
			UserCurrency: req.UserCurrency,
			Email:        req.Email,
			Consignee: model.Consignee{
				Email: req.Email,
			},
			Status: model.OrderStatusPending,
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.City = a.City
			o.Consignee.Country = a.Country
			o.Consignee.State = a.State
			o.Consignee.Street = a.StreetAddress
		}
		if err := tx.Create(&o).Error; err != nil {
			return err
		}
		items := make([]model.OrderItem, len(req.OrderItems))
		for _, oi := range req.OrderItems {
			items = append(items, model.OrderItem{
				ProductID: oi.Item.ProductId,
				OrderID:   orderId.String(),
				Quantity:  oi.Item.Quantity,
				Cost:      oi.Cost,
			})
		}

		if err := tx.Create(&items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	return
}
