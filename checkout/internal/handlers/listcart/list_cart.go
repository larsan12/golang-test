package listcart

import (
	"context"
	"log"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers"

	"github.com/pkg/errors"
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
	User int64 `json:"user"`
}

type Response struct {
	Items      []domain.CartItem `json:"items"`
	TotalPrice uint32            `json:"totalPrice"`
}

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("[handler listCart] %+v", req)
	var response Response
	// TODO
	items, err := h.businessLogic.ListCart(ctx, req.User)
	if err != nil {
		return response, errors.Wrap(err, "[handler listCart] error")
	}

	var totalPrice uint32
	for _, item := range items {
		totalPrice += item.Price
	}

	response.TotalPrice = totalPrice
	response.Items = items

	return response, nil
}
