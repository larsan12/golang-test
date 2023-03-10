package checkout_v1

import (
	"context"
	"route256/checkout/internal/domain"
	desc "route256/checkout/pkg/checkout_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemetation) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	err := i.businessLogic.DeleteFromCart(ctx, domain.CartItemDiff{
		User: req.GetUser(),
		Sku: req.GetSku(),
		Count: uint16(req.GetCount()),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
