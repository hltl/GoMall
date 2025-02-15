package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/hltl/GoMall/app/cart/utils"
	"github.com/hltl/GoMall/app/cart/conf"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)
}
