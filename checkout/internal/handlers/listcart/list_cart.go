package listcart

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

type ProductItem struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Response struct {
	Items      []ProductItem `json:"items"`
	TotalPrice uint32        `json:"totalPrice"`
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
	log.Printf("[handler] listCart: %+v", req)
	var response Response
	// TODO
	items, err := h.businessLogic.ListCart(ctx, req.User)
	responseItems := make([]ProductItem, 0, len(items))
	if err != nil {
		return response, errors.Wrap(err, "[handler] listCart error")
	}

	var totalPrice uint32
	for _, item := range items {
		totalPrice += item.Price
		responseItems = append(responseItems, ProductItem{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	response.TotalPrice = totalPrice
	response.Items = responseItems

	return response, nil
}
