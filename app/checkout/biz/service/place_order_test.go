package service

import (
	"context"
	"testing"
	checkout "github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	req := &checkout.PlaceOrderReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}