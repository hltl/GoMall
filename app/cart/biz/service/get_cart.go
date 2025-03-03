package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/cart/biz/model"
	"github.com/hltl/GoMall/app/cart/biz/dal/mysql"
	cart "github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	if req.UserId == 0{
		return nil,kerrors.NewGRPCBizStatusError(200401,"user id is required")
	}
	r, err := model.GetItemsByUserId(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil{
		return nil, err
	}
	resp = &cart.GetCartResp{Cart: &cart.Cart{}}
	resp.Cart.UserId = req.UserId
	for _, v:=range r{
		resp.Cart.Items=append(resp.Cart.Items, &cart.CartItem{ProductId: uint32(v.ProductId),Quantity: int32(v.Quantity)})
	}
	return resp,nil
}
