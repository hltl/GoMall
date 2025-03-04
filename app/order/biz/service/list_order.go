package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/order/biz/dal/mysql"
	"github.com/hltl/GoMall/app/order/biz/model"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
	order "github.com/hltl/GoMall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	if req.UserId == 0 {
		return nil, kerrors.NewGRPCBizStatusError(6004001, "user id is required")
	}

	orders, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(6005001, fmt.Sprintf("list order failed:%v", err))
	}
	resp = &order.ListOrderResp{}
	for _, v := range orders {
		items := make([]*order.OrderItem, 0)
		for _, it := range v.Items {
			items = append(items, &order.OrderItem{
				Cost: it.Cost,
				Item: &cart.CartItem{
					ProductId: it.ProductID,
					Quantity:  it.Quantity,
				},
			})
		}
		resp.Orders = append(resp.Orders, &order.Order{
			CreatedAt:    int32(v.CreatedAt.Unix()),
			OrderId:      v.OrderID,
			UserId:       v.UserID,
			UserCurrency: v.UserCurrency,
			Email:        v.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.Street,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		})
	}
	return
}
