package service

import (
	"context"

	"github.com/hltl/GoMall/app/cart/biz/dal/mysql"
	"github.com/hltl/GoMall/app/cart/biz/model"
	cart "github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	if req.UserId ==0{
		return 
	}
	err=model.EmptyCart(s.ctx,mysql.DB,req.UserId)
	return
}
