package service

import (
	"context"

	"github.com/hltl/GoMall/app/auth/biz/dal/cache"
	auth "github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	valid, err := cache.ValidateToken(s.ctx, uint(req.UserId),req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.VerifyResp{Res: valid}, nil
}

