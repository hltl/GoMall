package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/hltl/GoMall/gomall/app/frontend/utils"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/order"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	p, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(frontendutils.GetUserIdFromCtx(h.Context))})
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, 0)
	for _, v := range p.Orders {
		for _, item := range v.OrderItems {
			ids = append(ids, item.Item.ProductId)
		}
	}
	pds, err := rpc.ProductClient.GetProducts(h.Context, &product.GetProductsReq{Ids: ids})
	if err != nil {
		return nil, err
	}

	products := make(map[uint32]*product.Product)
	for _, v := range pds.Products {
		products[v.Id] = v
	}

	var orders []map[string]any
	for _, v := range p.Orders {
		var items []map[string]string
		for _, item := range v.OrderItems {
			items = append(items, map[string]string{
				"Picture":     products[item.Item.ProductId].Picture,
				"ProductName": products[item.Item.ProductId].Name,
				"Qty":         strconv.Itoa(int(item.Item.Quantity)),
				"Cost":        strconv.FormatFloat(float64(item.Cost), 'f', 2, 64),
			})
		}
		t := time.Unix(int64(v.CreatedAt), 0)
		orders = append(orders, map[string]any{
			"OrderId":    v.OrderId,
			"CreatedDate": t,
			"Items":      items,
		})
	}
	return utils.H{
		"orders": orders,
		"title":  "Order List",
	}, nil
}
