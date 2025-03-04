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
    if len(req.OrderItems) == 0 {
        return nil, kerrors.NewGRPCBizStatusError(6004001, "order items is required")
    }

    err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
        // 创建订单
        orderId, err := uuid.NewUUID()
        if err != nil {
            return kerrors.NewGRPCBizStatusError(6004002, "generate order id failed")
        }

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

        // 处理地址信息
        if req.Address != nil {
            o.Consignee.City = req.Address.City
            o.Consignee.Country = req.Address.Country
            o.Consignee.State = req.Address.State
            o.Consignee.Street = req.Address.StreetAddress
        }

        // 创建订单记录
        if err := tx.Create(o).Error; err != nil {
            return kerrors.NewGRPCBizStatusError(6004003, "create order failed")
        }

        // 创建订单项
        items := make([]model.OrderItem, 0, len(req.OrderItems))
        for _, oi := range req.OrderItems {
            if oi.Item == nil {
                return kerrors.NewGRPCBizStatusError(6004004, "invalid order item")
            }
            items = append(items, model.OrderItem{
                ProductID: oi.Item.ProductId,
                OrderID:   orderId.String(),
                Quantity:  oi.Item.Quantity,
                Cost:      oi.Cost,
            })
        }

        // 批量创建订单项
        if len(items) > 0 {
            if err := tx.Create(&items).Error; err != nil {
                return kerrors.NewGRPCBizStatusError(6004005, err.Error())
            }
        }

        resp = &order.PlaceOrderResp{
            Order: &order.OrderResult{
                OrderId: orderId.String(),
            },
        }
        return nil
    })

    if err != nil {
        return nil, err
    }
    return resp, nil
}
