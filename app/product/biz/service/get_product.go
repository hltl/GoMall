package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hltl/GoMall/app/product/biz/dal/mysql"
	"github.com/hltl/GoMall/app/product/biz/model"
	product "github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(200401, "product id is required")
	}
	p, err := model.GetProductById(s.ctx, mysql.DB, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Description: p.Description,
			Name:        p.Name,
			Price:       p.Price,
		},
	}, nil
}
