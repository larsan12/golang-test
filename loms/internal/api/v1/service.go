package loms_v1

import (
	"route256/loms/internal/domain"
	desc "route256/loms/pkg/loms_v1"
)

type Implemetation struct {
	desc.UnimplementedLomsV1Server
	businessLogic domain.BusinessLogic
}

func NewLomsV1(businessLogic domain.BusinessLogic) *Implemetation {
	return &Implemetation{
		desc.UnimplementedLomsV1Server{},
		businessLogic,
	}
}
