package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hltl/GoMall/gomall/app/frontend/biz/service"
	"github.com/hltl/GoMall/gomall/app/frontend/biz/utils"
	auth "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/auth"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusOK, []byte("/"))
}

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
}
