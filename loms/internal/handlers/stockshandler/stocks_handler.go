package stockshandler

import (
	"context"
	"log"
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type Request struct {
	SKU uint32 `json:"sku"`
}

func (r Request) Validate() error {
	// TODO: implement
	return nil
}

type Response struct {
	Stocks []domain.StockItem `json:"stocks"`
}

type Handler struct {
	businessLogic domain.BusinessLogic
}

func New(businessLogic domain.BusinessLogic) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("[handler stocks] %+v", request)
	var response Response
	stocks, err := h.businessLogic.Stock(ctx, request.SKU)
	if err != nil {
		return response, errors.WithMessage(err, "[handler stocks] Stock error")
	}
	response.Stocks = stocks
	return response, nil
}
