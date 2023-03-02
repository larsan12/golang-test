package checkout_v1

import (
	"context"
	desc "route256/checkout/pkg/checkout_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemetation) Puchase(ctx context.Context, req *desc.PuchaseRequest) (*emptypb.Empty, error) {
	_, err := i.businessLogic.Puchase(ctx, req.User)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
