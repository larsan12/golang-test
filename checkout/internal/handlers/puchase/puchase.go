package puchase

import (
	"context"
	"log"
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
	log.Printf("[handler puchase] %+v", req)
	var response Response
	order, err := h.businessLogic.Puchase(ctx, req.User)
	if err != nil {
		return response, errors.Wrap(err, "[handler puchase] Puchase error")
	}
	log.Printf("[handler puchase]  successfully puchase order with orderId: %d", order)

	return response, nil
}
