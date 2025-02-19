package middleware

import (
	"context"
	"fmt"
	"path"
	"runtime"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
)

type CasbinMiddleware struct {
	enforcer *casbin.Enforcer
}

func NewCasbinMiddleware(modelPath, policyPath string) (*CasbinMiddleware, error) {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  path.Base(file),
			"line":  line,
			"error": err,
		}).Error("初始化Casbin失败")
		return nil, fmt.Errorf("初始化Casbin失败: %v", err)
	}

	// 加载策略
	if err := enforcer.LoadPolicy(); err != nil {
		logrus.WithError(err).Error("加载Casbin策略失败")
		return nil, fmt.Errorf("加载Casbin策略失败: %v", err)
	}

	return &CasbinMiddleware{enforcer: enforcer}, nil
}

func (cm *CasbinMiddleware) AuthorizeMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从上下文获取用户角色
		role := c.GetString("user_role")
		if role == "" {
			role = "anonymous" // 默认角色
		}

		// 获取请求路径和方法
		path := string(c.Request.URI().Path())
		method := string(c.Request.Method())

		// 检查权限
		allowed, err := cm.enforcer.Enforce(role, path, method)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"role":   role,
				"path":   path,
				"method": method,
				"error":  err,
			}).Error("Casbin权限检查失败")
			c.AbortWithStatus(500)
			return
		}

		if !allowed {
			logrus.WithFields(logrus.Fields{
				"role":   role,
				"path":   path,
				"method": method,
			}).Warn("访问被拒绝")
			c.AbortWithStatus(403)
			return
		}

		logrus.WithFields(logrus.Fields{
			"role":   role,
			"path":   path,
			"method": method,
		}).Debug("权限检查通过")
		c.Next(ctx)
	}
}
