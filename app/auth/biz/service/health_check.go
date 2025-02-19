package service

import (
	"context"
	auth "github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
)

type HealthCheckService struct {
	ctx context.Context
} // NewHealthCheckService new HealthCheckService
func NewHealthCheckService(ctx context.Context) *HealthCheckService {
	return &HealthCheckService{ctx: ctx}
}

// Run create note info
func (s *HealthCheckService) Run(req *auth.HealthCheckReq) (resp *auth.HealthCheckResp, err error) {
	// Finish your business logic.

	return
}
