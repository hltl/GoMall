package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/hltl/GoMall/gomall/app/frontend/utils"
	rpccart "github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	c, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var total float64

	var ids []uint32

	for _, v := range c.Cart.Items {
		ids = append(ids, v.ProductId)
	}

	pds, err := rpc.ProductClient.GetProducts(h.Context, &rpcproduct.GetProductsReq{Ids: ids})
	if err != nil {
		return nil, err
	}

	for i, v := range pds.Products {
		items = append(items, map[string]string{
			"Name":    v.Name,
			"Price":   strconv.FormatFloat(float64(v.Price), 'f', 2, 64),
			"Picture": v.Picture,
			"Qty":     strconv.Itoa(int(c.Cart.Items[i].Quantity)),
		})
		total += float64(c.Cart.Items[i].Quantity) * float64(v.Price)
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
