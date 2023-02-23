package deletefromcart

import (
	"context"
	"log"
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic domain.BusinessLogic
}

func New(businessLogic domain.BusinessLogic) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

type Request struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

var (
	ErrEmptyUser = errors.New("empty user")
	ErrEmptySKU  = errors.New("empty sku")
)

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	if r.Sku == 0 {
		return ErrEmptySKU
	}
	return nil
}

type Response struct{}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("[handler deleteFromCart] %+v", req)
	var response Response
	error := h.businessLogic.DeleteFromCart(ctx, req.User, req.Sku, req.Count)
	if error != nil {
		return response, errors.WithMessage(error, "[handler] deleteFromCart deletion error")
	}
	return response, nil
}
