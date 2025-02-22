package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
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

	resp = make(map[string]any)
	resp["Title"]="Home"
	items := []map[string]any{
		{"Name":"test1","Price":100,"Picture":"/static/image/test1.jpg"},
		{"Name":"test2","Price":200,"Picture":"/static/image/test1.jpg"},
		{"Name":"test3","Price":300,"Picture":"/static/image/test1.jpg"},
		{"Name":"test1","Price":100,"Picture":"/static/image/test1.jpg"},
		{"Name":"test2","Price":200,"Picture":"/static/image/test1.jpg"},
		{"Name":"test3","Price":300,"Picture":"/static/image/test1.jpg"},
	}
	resp["Items"]=items
	// todo edit your code
	return resp, nil
}
