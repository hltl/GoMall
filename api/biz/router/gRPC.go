package router

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hltl/GoMall/api/biz/proto/auth" // 替换为实际路径

	"google.golang.org/grpc"
)

// AuthService 实现
type AuthService struct {
	auth.UnimplementedAuthServiceServer
	// 可添加数据库连接等依赖项
}

// DeliverTokenByRPC 实现
func (s *AuthService) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (*auth.DeliveryResp, error) {
	// 实际生成token的逻辑
	token := generateJWT(req.UserId) // 示例函数
	return &auth.DeliveryResp{
		Token: token,
	}, nil
}

// VerifyTokenByRPC 实现
func (s *AuthService) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (*auth.VerifyResp, error) {
	// 实际验证逻辑
	isValid := validateJWT(req.Token) // 示例函数
	return &auth.VerifyResp{
		Res: isValid,
	}, nil
}

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ConnectionTimeout(5 * time.Second),
	)
	auth.RegisterAuthServiceServer(s, &AuthService{})

	// 优雅关闭
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down gRPC server...")
		s.GracefulStop()
	}()

	log.Printf("gRPC server listening on :%s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// 示例JWT生成函数
func generateJWT(userID int32) string {
	// 实际实现需要包含签名、过期时间等
	return fmt.Sprintf("jwt.token.%d", userID)
}

// 示例验证函数
func validateJWT(token string) bool {
	// 实际需要解析验证JWT
	return token != ""
}
