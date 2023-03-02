package checkout_v1

import (
	"route256/checkout/internal/domain"
	desc "route256/checkout/pkg/checkout_v1"
)

type Implemetation struct {
	desc.UnimplementedCheckoutV1Server
	businessLogic domain.BusinessLogic
}

func NewCheckoutV1(businessLogic domain.BusinessLogic) *Implemetation {
	return &Implemetation{
		desc.UnimplementedCheckoutV1Server{},
		businessLogic,
	}
}
