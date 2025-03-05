package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/auth/biz/dal/cache"
	"github.com/hltl/GoMall/app/auth/utils"
	auth "github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
	"github.com/redis/go-redis/v9"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	tok, err := cache.GetUserToken(s.ctx, uint(req.UserId))
	if err == redis.Nil {
		// Token 不存在
        tok,err= utils.GenerateToken(uint(req.UserId))
        println("tok",tok)
		if err != nil {
            return nil, kerrors.NewGRPCBizStatusError(1005001, "Failed to generate token")
        }
        err = cache.SetUserToken(s.ctx, uint(req.UserId), tok)
        
        if err != nil {
            println("err",err.Error())
            return nil, kerrors.NewGRPCBizStatusError(1005002, "Failed to set token")
        }
        
	} else if err != nil {
		// 查询 Redis 时发生错误
		return nil, kerrors.NewGRPCBizStatusError(1005001, fmt.Sprintf("Failed to get token from Redis: %v", err))
	}

	// Token 存在，返回 token
	return &auth.DeliveryResp{Token: tok}, nil
}
