package orderpayedhandler

import (
	"context"
	"log"
	"route256/loms/internal/handlers"

	"github.com/pkg/errors"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

type Response struct{}

func (r Request) Validate() error {
	// TODO: implement
	return nil
}

type Handler struct {
	businessLogic handlers.BusinessLogic
}

func New(businessLogic handlers.BusinessLogic) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("[handler orderPayed] %+v", request)
	var response Response
	err := h.businessLogic.OrderPayed(ctx, request.OrderID)
	if err != nil {
		return response, errors.WithMessage(err, "[handler orderPayed] error")
	}
	return response, nil
}
