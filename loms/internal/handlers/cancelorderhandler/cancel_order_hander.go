package cancelorderhandler

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
	log.Printf("[handler cancelOrder] %+v", request)
	var response Response
	err := h.businessLogic.CancelOrder(ctx, request.OrderID)
	if err != nil {
		return response, errors.WithMessage(err, "[handler cancelOrder] error")
	}
	return response, nil
}
