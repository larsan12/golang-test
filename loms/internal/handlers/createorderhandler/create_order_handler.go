package createorderhandler

import (
	"context"
	"log"
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type Request struct {
	User  int64              `json:"user"`
	Items []domain.OrderItem `json:"items"`
}

func (r Request) Validate() error {
	// TODO: implement
	return nil
}

type Response struct {
	OrderID int64 `json:"orderID"`
}

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("[handler createOrder] %+v", request)
	var response Response
	orderId, err := h.businessLogic.CreateOrder(ctx, request.User, request.Items)
	if err != nil {
		return response, errors.WithMessage(err, "[handler createOrder] error")
	}
	response.OrderID = orderId
	return response, nil

}
