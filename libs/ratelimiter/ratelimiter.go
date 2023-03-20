package ratelimiter

import (
	"context"
	"log"
	"time"
)

// описание внешнего интерфейса

type Limiter interface {
	Wait(ctx context.Context)
	Close()
}

// проверка типов
var _ Limiter = implemetation{}

type implemetation struct {
	rps     uint
	burst   uint
	limiter chan struct{}
	stop    chan struct{}
}

// rps - лимит запросов в секунду
// burst - колличество одновременных запросов
func NewLimiter(rps uint, burst uint) Limiter {

	if rps > 1_000_000_000 {
		log.Fatal("rps must be less than 10^9")
	}

	if burst > rps {
		log.Fatal("burst must be less or equal rps")
	}

	// создаем канал с буфером размером {burst}
	limiter := make(chan struct{}, burst)

	// канал для завершения лимитера
	stop := make(chan struct{})

	// одна запись в канал равно одной квоте на запрос, заполняем буфер тем самым позволяя сделать сразу {burst} запросов одновременно
	for i := 0; i < int(burst); i++ {
		limiter <- struct{}{}
	}

	// считаем частоту запросов - сколько раз в наносекунду создавать квоту на запрос
	frequency := 1_000_000_000.0 / rps

	// горутина создающая квоту на запрос раз в {frequency} наносекунд
	go func() {
		for range time.Tick(time.Duration(frequency) * time.Nanosecond) {
			select {
			// закрываем горутину в случае остановки
			case <-stop:
				return
			default:
				limiter <- struct{}{}
			}
		}
	}()

	return implemetation{rps, burst, limiter, stop}
}

// при запросе ждём квоту из канала limiter
// слушаем также завершение контекста
func (l implemetation) Wait(ctx context.Context) {
	select {
	case <-ctx.Done():
	case <-l.limiter:
	}
}

func (l implemetation) Close() {
	close(l.stop)
	close(l.limiter)
}
