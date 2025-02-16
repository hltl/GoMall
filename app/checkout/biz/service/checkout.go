package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hltl/GoMall/app/checkout/rpc"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/cart"
	checkout "github.com/hltl/GoMall/rpc_gen/kitex_gen/checkout"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/order"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/payment"
	"github.com/hltl/GoMall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, fmt.Sprintf("get cart failed:%v", err))
	}
	if cartResp.Cart == nil || cartResp.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(50054001, "cart is empty")
	}
	// 获取购物车中的商品信息
	productsReq := &product.GetProductsReq{}
	for _, item := range cartResp.Cart.Items {
		productsReq.Ids = append(productsReq.Ids, item.ProductId)
	}
	productsResp, err := rpc.ProductClient.GetProducts(s.ctx, productsReq)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005002, fmt.Sprintf("get products failed:%v", err))
	}
	if len(productsResp.Products) < len(cartResp.Cart.Items) {
		return nil, kerrors.NewGRPCBizStatusError(50054002, "get products failed")
	}
	// 计算总价,生成订单条目
	price := make(map[uint32]float32, len(productsResp.Products))
	for _, p := range productsResp.Products {
		price[p.Id] = p.Price
	}
	var total float32
	orderItem := make([]*order.OrderItem, len(productsResp.Products))
	for _, item := range cartResp.Cart.Items {
		cost := price[item.ProductId] * float32(item.Quantity)
		total += cost
		orderItem = append(orderItem, &order.OrderItem{Item: item, Cost: cost})
	}

	// 创建订单
	var orderId string
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD", // 简化处理，实际应该根据用户信息获取
		Email:        req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		OrderItems: orderItem,
	}
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		return nil, err
	}
	orderId = orderResp.Order.OrderId
	payReq := &payment.ChargeReq{
		OrderId: orderId,
		UserId:  req.UserId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}

	// 交易支付

	payResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005003, fmt.Sprintf("payment failed:%v", err))
	}
	if _, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId}); err != nil {
		klog.Error("mark order failed:%v",err)
		// todo 数据不一致
	}

	klog.Info(payResp)
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error("empty cart failed:%v", err)
	}
	return &checkout.CheckoutResp{OrderId: orderId, TransactionId: payResp.TransactionId}, nil
}
