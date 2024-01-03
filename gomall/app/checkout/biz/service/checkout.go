package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/baiyutang/gomall/app/checkout/infra/rpc"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/cart"
	checkout "github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/order"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/payment"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

/*
	Run

1. get cart
2. calculate cart
3. create order
4. empty cart
5. pay
6. change order result
7. finish
*/
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutRes, err error) {
	// Finish your business logic.
	// Idempotent
	// get cart
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartRequest{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}
	if cartResult == nil || len(cartResult.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Items {
		p, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductRequest{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	// create order
	orderReq := &order.PlaceOrderRequest{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)
	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartRequest{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)
	// charge
	var orderId string
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	payReq := &payment.ChargeRequest{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		}}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}

	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)

	resp = &checkout.CheckoutRes{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}