package checkout

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hltl/GoMall/gomall/app/frontend/biz/service"
	"github.com/hltl/GoMall/gomall/app/frontend/biz/utils"
	checkout "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/checkout"
	common "github.com/hltl/GoMall/gomall/app/frontend/hertz_gen/frontend/common"
)

// Checkout .
// @router /checkout [GET]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "checkout", utils.WrapResponse(ctx, c, resp))
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "waiting", utils.WrapResponse(ctx, c, resp))
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "result", utils.WrapResponse(ctx, c, resp))
}
