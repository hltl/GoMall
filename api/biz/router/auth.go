package router

import (
	"context"
	"fmt"
	"time"

	"github.com/hltl/GoMall/api/biz/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	conn    *grpc.ClientConn
	Service auth.AuthServiceClient
}

func NewAuthClient(serverAddr string) (*AuthClient, error) {
	// 建立gRPC连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("连接认证服务失败: %v", err)
	}

	return &AuthClient{
		conn:    conn,
		Service: auth.NewAuthServiceClient(conn),
	}, nil
}

// 生成Token
func (c *AuthClient) DeliverToken(userID int32) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.Service.DeliverTokenByRPC(ctx, &auth.DeliverTokenReq{
		UserId: userID,
	})
	if err != nil {
		return "", fmt.Errorf("生成Token失败: %v", err)
	}

	return resp.Token, nil
}

// 验证Token
func (c *AuthClient) VerifyToken(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.Service.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		return false, fmt.Errorf("验证Token失败: %v", err)
	}

	return resp.Res, nil
}

// 关闭连接
func (c *AuthClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (s *AuthService) HealthCheck(ctx context.Context, req *auth.HealthCheckReq) (*auth.HealthCheckResp, error) {
	return &auth.HealthCheckResp{Status: "OK"}, nil
}
