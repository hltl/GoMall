package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/cart/biz/dal/mysql"
	"github.com/hltl/GoMall/app/cart/biz/model"
	"github.com/hltl/GoMall/app/cart/rpc"
	cart "github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
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
	presp,err:=rpc.ProductClient.GetProduct(s.ctx,&product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil,err
	}
	if presp.Product == nil || presp.Product.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(200402, "product not found")
	}
	if req.UserId == 0 {
		return nil, kerrors.NewGRPCBizStatusError(200401, "user id is required")
	}
	err = model.AddItem(s.ctx, mysql.DB, req.UserId, model.Item{ProductId: req.Item.ProductId, Quantity: req.Item.Quantity})
	return
}
