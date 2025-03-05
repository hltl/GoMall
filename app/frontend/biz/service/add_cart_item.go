package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	cart "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/hltl/GoMall/gomall/app/frontend/utils"
	rpccart "github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "AddItem req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	_,err=rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{UserId: uint32(frontendutils.GetUserIdFromCtx(h.Context)), 
		Item:&rpccart.CartItem{ProductId: req.ProductId,Quantity: req.ProductNum}})
	if err!=nil{
		return nil,err
	}
	return
}
