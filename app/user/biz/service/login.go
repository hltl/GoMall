package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/user/biz/dal/mysql"
	"github.com/hltl/GoMall/app/user/biz/model"
	"github.com/hltl/GoMall/app/user/infra/rpc"
	"github.com/hltl/GoMall/app/user/utils"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
	user "github.com/hltl/GoMall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if !utils.CheckEmail(req.Email) {
		return nil, kerrors.NewGRPCBizStatusError(1004001, fmt.Sprintf("邮箱格式不正确:%s", req.Email))
	}
	u, err := model.GetUserByEmail(s.ctx, mysql.DB, req.Email)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(1004001, "用户不存在")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, kerrors.NewGRPCBizStatusError(1004001, "密码错误")
	}
	tok, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: int32(u.ID)})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(1005001, err.Error())
	}
	resp = &user.LoginResp{
		UserId: int32(u.ID),
		Token:  tok.Token,
	}
	return
}
