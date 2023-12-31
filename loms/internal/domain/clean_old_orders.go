package domain

import (
	"context"
	"fmt"
	"route256/libs/workerpool"
	"time"
)

var isRunning bool = false

const observationPeriod int = 10 // in minutes

func (m *Model) ObserveOldOrders(ctx context.Context, orderExpirationTime int) {
	if isRunning {
		return
	}

	go func() {
		// тикер каждые 10 минут
		for range time.Tick(time.Duration(observationPeriod) * time.Minute) {
			// берём заказы старше 10 минут
			orders, err := m.repository.GetOldOrders(ctx, time.Now().Add(-time.Duration(orderExpirationTime)*time.Minute))
			if err != nil {
				fmt.Println("Clean all orders, undexpected error: ", err)
				continue
			}
			fmt.Println("Clean all orders, count: ", len(orders))

			// таска
			cleanOldOrder := func(order Order) (bool, error) {
				err := m.CancelOrder(ctx, order.OrderId)
				return true, err
			}

			// create tasks
			tasks := make([]workerpool.Task[Order, bool], len(orders))
			for i, order := range orders {
				tasks[i] = workerpool.Task[Order, bool]{
					Run:    cleanOldOrder,
					InArgs: order,
				}
			}

			// выполняем в пуле
			_, err = m.orderCleanerWorkerPool.Execute(ctx, tasks)
			if err != nil {
				fmt.Println("Clean all orders, undexpected pool error: ", err)
				continue
			}

		}
		isRunning = false
	}()

	isRunning = true
}
