package main

import (
	"log"
	"net/http"
	"route256/libs/srvwrapper"
	"route256/loms/internal/handlers/cancelorderhandler"
	"route256/loms/internal/handlers/createorderhandler"
	"route256/loms/internal/handlers/listorderhandler"
	"route256/loms/internal/handlers/orderpayedhandler"
	"route256/loms/internal/handlers/stockshandler"
)

const port = ":8081"

func main() {

	stocksHandler := stockshandler.New()
	orderHandler := createorderhandler.New()
	listOrderHandler := listorderhandler.New()
	orderPayedHandler := orderpayedhandler.New()
	cancelOrderHandler := cancelorderhandler.New()

	http.Handle("/stocks", srvwrapper.New(stocksHandler.Handle))
	http.Handle("/createOrder", srvwrapper.New(orderHandler.Handle))
	http.Handle("/listOrder", srvwrapper.New(listOrderHandler.Handle))
	http.Handle("/orderPayed", srvwrapper.New(orderPayedHandler.Handle))
	http.Handle("/cancelOrder", srvwrapper.New(cancelOrderHandler.Handle))

	log.Println("listening http at", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}
