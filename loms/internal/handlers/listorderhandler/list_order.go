package listorderhandler

import (
	"context"
	"log"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers"

	"github.com/pkg/errors"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

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

func (h *Handler) Handle(ctx context.Context, request Request) (domain.Order, error) {
	log.Printf("[handler listOrder] %+v", request)
	order, err := h.businessLogic.ListOrder(ctx, request.OrderID)
	if err != nil {
		return order, errors.WithMessage(err, "[handler listOrder] ListOrder error")
	}
	return order, nil
}
