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

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterRequest) (redirect string, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "redirect = %+v", redirect)
	}()
	resp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email: req.Email, Password: req.Password, ConfirmPassword: req.PasswordConfirm,
	})
	if err!=nil{
		return "",err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err=session.Save()

	if err != nil {
		return "", err
	}
	redirect = "/"
	// todo edit your code
	return
}
