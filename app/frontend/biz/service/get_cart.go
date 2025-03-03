package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/hltl/GoMall/gomall/app/frontend/utils"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	defer func() {
	hlog.CtxInfof(h.Context, "req = %+v", req)
	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	p,err:=rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: uint32(frontendutils.GetUserIdFromCtx(h.Context))})
	if err!=nil{
		hlog.Error(h.Context, "GetCart failed", uint32(frontendutils.GetUserIdFromCtx(h.Context)))
		return nil,err
	}

	ids:=make([]uint32,len(p.Cart.Items))
	for i,v:=range p.Cart.Items{
		ids[i]=v.ProductId
	}
	pds,err:= rpc.ProductClient.GetProducts(h.Context, &product.GetProductsReq{Ids: ids})
	if err!=nil{
		return nil,err
	}
	items:=make([]map[string]string,len(pds.Products))
	var total float64
	for i,v:=range pds.Products{
		items[i]=map[string]string{
			"Name":v.Name,
			"Description":v.Description,
			"Price":strconv.FormatFloat(float64(v.Price), 'f', 2, 64),
			"Quantity":strconv.Itoa(int(p.Cart.Items[i].Quantity)),
		}
		total+=float64(v.Price)*float64(p.Cart.Items[i].Quantity)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": total,
	},nil
}
