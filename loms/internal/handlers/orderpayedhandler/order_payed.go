package orderpayedhandler

import (
	"context"
	"log"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

type Response struct{}

func (r Request) Validate() error {
	// TODO: implement
	return nil
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("[handler] orderPayed: %+v", request)
	return Response{}, nil
}
