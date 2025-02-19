package main

import (
	"context"
	auth "github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
	"github.com/hltl/GoMall/app/auth/biz/service"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}

// HealthCheck implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) HealthCheck(ctx context.Context, req *auth.HealthCheckReq) (resp *auth.HealthCheckResp, err error) {
	resp, err = service.NewHealthCheckService(ctx).Run(req)

	return resp, err
}
