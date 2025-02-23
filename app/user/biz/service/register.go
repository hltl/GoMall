package service

import (
	"context"

	"github.com/hltl/GoMall/app/user/biz/dal/mysql"
	"github.com/hltl/GoMall/app/user/biz/model"
	"github.com/hltl/GoMall/app/user/model"
	"github.com/hltl/GoMall/app/user/utils"
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
		return
	}
	// 校验邮箱
	if !utils.CheckEmail(req.Email) {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "邮箱格式不正确",
		}
		return
	}
	// 校验密码
	if !utils.CheckPassword(req.Password) {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "密码格式不正确",
		}
		return
	}
	// 检查邮箱是否已经注册
	if _,err =model.GetUserByEmail(s.ctx,mysql.DB,req.Email); err == nil {
		resp = &user.RegisterResp{
			Code:    1,
			Message: "邮箱已注册",
		}
		return
	}
	// 密码加密
	passwardHashed,err:=bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	if err!=nil{
		resp = &user.RegisterResp{
			Code:    1,
			Message: "密码加密失败",
		}
		return
	}
	u := &model.User{
		Email:    req.Email,
		Password: string(passwardHashed),
	}
	if err = model.CreateUser(s.ctx,mysql.DB,u);err!=nil{
		resp =&user.RegisterResp{
			Code:1,
			Message:"注册失败",
		}
		return
	}
	resp = &user.RegisterResp{
		Code:    0,
		Message: "注册成功",
		Token:,
	}

}
