package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/hltl/GoMall/gomall/app/frontend/conf"
	"github.com/hltl/GoMall/gomall/app/frontend/utils"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	CartClient cartservice.Client
	CheckoutClient checkoutservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClinet()
		initCartClinet()
		initCheckoutClinet()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Hertz.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	utils.MustHandleError(err)
}

func initProductClinet(){
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Hertz.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	utils.MustHandleError(err)
}

func initCartClinet(){
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Hertz.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	utils.MustHandleError(err)
}

func initCheckoutClinet(){
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Hertz.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	utils.MustHandleError(err)
}