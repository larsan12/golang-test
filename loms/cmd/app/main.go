package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/config"
	"route256/loms/internal/domain"
	"route256/loms/internal/handlers/cancelorderhandler"
	"route256/loms/internal/handlers/createorderhandler"
	"route256/loms/internal/handlers/listorderhandler"
	"route256/loms/internal/handlers/orderpayedhandler"
	"route256/loms/internal/handlers/stockshandler"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	businessLogic := domain.New()
	stocksHandler := stockshandler.New(businessLogic)
	orderHandler := createorderhandler.New(businessLogic)
	listOrderHandler := listorderhandler.New(businessLogic)
	orderPayedHandler := orderpayedhandler.New(businessLogic)
	cancelOrderHandler := cancelorderhandler.New(businessLogic)

	http.Handle("/stocks", srvwrapper.New(stocksHandler.Handle))
	http.Handle("/createOrder", srvwrapper.New(orderHandler.Handle))
	http.Handle("/listOrder", srvwrapper.New(listOrderHandler.Handle))
	http.Handle("/orderPayed", srvwrapper.New(orderPayedHandler.Handle))
	http.Handle("/cancelOrder", srvwrapper.New(cancelOrderHandler.Handle))

	log.Println("listening http at", config.ConfigData.Port)
	err = http.ListenAndServe(":" + config.ConfigData.Port, nil)
	log.Fatal("cannot listen http", err)
}
