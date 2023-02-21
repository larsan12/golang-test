package createorderhandler

import (
	"context"
	"log"
)

type Request struct {
	User  int64  `json:"user"`
	Items []Item `json:"items"`
}

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (r Request) Validate() error {
	// TODO: implement
	return nil
}

type Response struct {
	OrderID int64 `json:"orderID"`
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("[handler] createOrder: %+v", request)
	return Response{
		OrderID: 10,
	}, nil
}
