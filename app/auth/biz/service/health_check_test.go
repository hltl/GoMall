package service

import (
	"context"
	"testing"
	auth "github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
)

func TestHealthCheck_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHealthCheckService(ctx)
	// init req and assert value

	req := &auth.HealthCheckReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
