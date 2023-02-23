package addtocart

import (
	"context"
	"errors"
	"log"
	"route256/checkout/internal/handlers"
)

type Handler struct {
	businessLogic handlers.BusinessLogic
}

func New(businessLogic handlers.BusinessLogic) *Handler {
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
	log.Printf("[handler addToCart] %+v", req)

	var response Response = Response{}

	err := h.businessLogic.AddToCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}
