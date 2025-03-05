package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/user/biz/dal/mysql"
	"github.com/hltl/GoMall/app/user/biz/model"
	"github.com/hltl/GoMall/app/user/infra/rpc"
	"github.com/hltl/GoMall/app/user/utils"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/auth"
	user "github.com/hltl/GoMall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Password != req.ConfirmPassword {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "两次密码输入不一致",
		}
		return resp, kerrors.NewGRPCBizStatusError(1004001, "两次密码输入不一致")
	}
	// 校验邮箱
	if !utils.CheckEmail(req.Email) {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "邮箱格式不正确",
		}
		return resp, kerrors.NewGRPCBizStatusError(1004001, "邮箱格式不正确")
	}
	// 检查邮箱是否已经注册
	if _, err = model.GetUserByEmail(s.ctx, mysql.DB, req.Email); err == nil {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "邮箱已注册",
		}
		return resp, kerrors.NewGRPCBizStatusError(1004001, "邮箱已注册")
	}
	// 密码加密
	passwardHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "密码加密失败",
		}
		return resp, kerrors.NewGRPCBizStatusError(1005001, "密码加密失败")
	}
	u := &model.User{
		Email:    req.Email,
		Password: string(passwardHashed),
	}
	if err = model.CreateUser(s.ctx, mysql.DB, u); err != nil {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "注册失败",
		}
		return resp, kerrors.NewGRPCBizStatusError(1005001, "注册失败")
	}
	tok, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: int32(u.ID)})
	if err != nil {
		_ = model.DeleteUser(s.ctx, mysql.DB, u.ID)
		return nil, kerrors.NewGRPCBizStatusError(1005001, err.Error())
	}
	resp = &user.RegisterResp{
		UserId:  int32(u.ID),
		Code:    0,
		Message: "注册成功",
		Token:   tok.Token,
	}
	return
}
