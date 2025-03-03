package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hltl/GoMall/gomall/app/frontend/infra/rpc"
	"github.com/hltl/GoMall/gomall/app/frontend/utils"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WrapResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	user_id := utils.GetUserIdFromCtx(ctx)
	content["user_id"] = user_id
	if user_id > 0 {
		p, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: uint32(user_id)})
		if err == nil && p != nil {
			content["cart_num"] = len(p.Cart.Items)
		}
	}
	return content
}
