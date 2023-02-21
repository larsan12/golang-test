package listorderhandler

import (
	"context"
	"log"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	Status string `json:"status"`
	User   int64  `json:"user"`
	Items  []Item `json:"items"`
}

func (r Request) Validate() error {
	// TODO: implement
	return nil
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("[handler] listOrder: %+v", request)

	order := [2]int32{1076963, 1148162}
	items := make([]Item, 0, len(order))
	for _, sku := range order {
		items = append(items, Item{
			Sku:   uint32(sku),
			Count: 1,
		})
	}

	return Response{
		Status: "new",
		User:   1,
		Items:  items,
	}, nil
}
