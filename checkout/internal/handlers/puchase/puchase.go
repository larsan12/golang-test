package puchase

import (
	"context"
	"log"
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

type Request struct {
	User int64 `json:"user"`
}

type Response struct{}

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
	log.Printf("[handler] puchase: %+v", req)
	var response Response
	cartItems, err := h.businessLogic.ListCart(ctx, req.User)
	if err != nil {
		return response, errors.Wrap(err, "[handler] puchase, ListCart error")
	}
	log.Printf("[handler] puchase, cartItems: %+v", cartItems)

	orderItems := make([]domain.OrderItem, 0, len(cartItems))

	for _, item := range cartItems {
		orderItems = append(orderItems, domain.OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}

	log.Printf("[handler] puchase, orderItems: %+v", orderItems)

	// TODO
	order, error := h.businessLogic.CreateOrder(ctx, req.User, orderItems)
	if error != nil {
		return response, errors.Wrap(err, "Puchase error")
	}
	log.Printf("[handler] success puchase, orderId: %d", order)

	return response, nil
}
