package auth

import (
	"context"
	auth "github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (resp *auth.VerifyResp, err error) {
	resp, err = defaultClient.VerifyTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func HealthCheck(ctx context.Context, req *auth.HealthCheckReq, callOptions ...callopt.Option) (resp *auth.HealthCheckResp, err error) {
	resp, err = defaultClient.HealthCheck(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "HealthCheck call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
