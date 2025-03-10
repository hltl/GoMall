package service

import (
	"context"

	"github.com/hltl/GoMall/app/product/biz/dal/mysql"
	"github.com/hltl/GoMall/app/product/biz/model"
	product "github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	products, err := model.SearchProduct(s.ctx, mysql.DB, req.Query)
	if err != nil {
		return nil, err
	}
	var r []*product.Product
	for _, v := range products {
		r = append(r, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}

	return &product.SearchProductsResp{Results: r}, nil
}
