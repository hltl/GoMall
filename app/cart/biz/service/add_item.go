package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/cart/biz/dal/mysql"
	"github.com/hltl/GoMall/app/cart/biz/model"
	cart "github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	if req.UserId == 0 {
		return nil, kerrors.NewGRPCBizStatusError(200401, "user id is required")
	}
	err = model.AddItem(s.ctx, mysql.DB, uint(req.UserId), model.Item{ProductId: uint(req.Item.ProductId), Quantity: int(req.Item.Quantity)})
	return
}
