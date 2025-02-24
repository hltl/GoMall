package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	auth "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginRequest) (redirect string, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "redirect = %+v", redirect)
	}()
	// todo edit your code
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		redirect = "/sign-in"
		return
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err = session.Save()
	if err != nil {
		redirect = "/sign-in"
		return
	}
	if req.Next != "" {
		redirect = req.Next
	} else {
		redirect = "/"
	}

	return
}
