package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (redirect string, err error) {
	defer func() {
	hlog.CtxInfof(h.Context, "req = %+v", req)
	hlog.CtxInfof(h.Context, "redirect = %+v", redirect)
	}()
	// todo edit your code

	s:=sessions.Default(h.RequestContext)
	s.Clear()
	s.Delete("user_id")
	if err= s.Save();err!=nil{
		redirect = "/"
		return 
	}
	redirect = "/sign-in"
	return
}
