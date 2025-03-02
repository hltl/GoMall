package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	defer func() {
	hlog.CtxInfof(h.Context, "req = %+v", req)
	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	p,err:=rpc.ProductClient.ListProducts(h.Context,&product.ListProductsReq{})
	if err!=nil{
		return nil,err
	}
	return utils.H{
		"title":"Hot sale",
		"items":p.Products,
	},nil
}
