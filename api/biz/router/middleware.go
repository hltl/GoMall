package router

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(authClient *AuthClient) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(401, map[string]string{"error": "未提供认证信息"})
			return
		}

		token := strings.TrimPrefix(string(authHeader), "Bearer ")
		if valid, err := authClient.VerifyToken(token); !valid || err != nil {
			logrus.WithFields(logrus.Fields{
				"token":     token,
				"client_ip": ctx.ClientIP(),
			}).Warn("无效的访问令牌")
			ctx.AbortWithStatusJSON(401, map[string]string{"error": "无效的访问令牌"})
			return
		}

		ctx.Next(c)
	}
}
