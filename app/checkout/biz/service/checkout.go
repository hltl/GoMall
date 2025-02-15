package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/checkout/rpc"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
	checkout "github.com/hltl/GoMall/rpc_gen/kitex_gen/checkout"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp,err:=rpc.CartClient.GetCart(s.ctx,&cart.GetCartReq{UserId:req.UserId})
	if err!=nil{
		return nil,kerrors.NewGRPCBizStatusError(5005001,fmt.Sprintf("get cart failed:%v",err))
	}
	if cartResp.Cart==nil || cartResp.Cart.Items==nil{
		return nil,kerrors.NewGRPCBizStatusError(50054001,"cart is empty")
	}
	// 获取购物车中的商品信息

	// 计算总价
	return
}
