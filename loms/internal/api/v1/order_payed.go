package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemetation) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	err := i.businessLogic.OrderPayed(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
