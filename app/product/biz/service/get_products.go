package service

import (
	"context"

	"github.com/hltl/GoMall/app/product/biz/dal/mysql"
	"github.com/hltl/GoMall/app/product/biz/model"
	product "github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type GetProductsService struct {
	ctx context.Context
} // NewGetProductsService new GetProductsService
func NewGetProductsService(ctx context.Context) *GetProductsService {
	return &GetProductsService{ctx: ctx}
}

// Run create note info
func (s *GetProductsService) Run(req *product.GetProductsReq) (resp *product.GetProductsResp, err error) {
	// Finish your business logic.
	// Example implementation: return products for even IDs and record odd IDs as failures.
	var failedIds []uint32
	products := make([]*product.Product, 0, len(req.Ids))
	// Assume req.ProductIds contains the list of product IDs.
	for _, id := range req.Ids {
		p,err:= model.GetProductById(s.ctx,mysql.DB,uint(id))
		if err!=nil || id == 0{
			failedIds = append(failedIds, id)
		}else{
			products = append(products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Picture:    p.Picture,
			})
		}

	}
	return &product.GetProductsResp{Products: products, Faileds: failedIds}, nil
}
