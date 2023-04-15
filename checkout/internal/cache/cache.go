package cache

import (
	"route256/libs/cache"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func Create[T any](ttl int) cache.Cache[T] {
	statReceiver := NewStatReceiver()
	return cache.New[T](ttl, statReceiver)
}

func NewStatReceiver() cache.StatsReceiver {
	var hitCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_cache_hits_total",
	})
	var requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_cache_requests_total",
	})
	var errorsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_cache_errors_total",
	})
	var settingTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_cache_seting_total",
	})
	return &statsReceiver{
		hitCount,
		requestsTotal,
		errorsTotal,
		settingTotal,
	}
}

type statsReceiver struct {
	hitCount      prometheus.Counter
	requestsTotal prometheus.Counter
	errorsTotal   prometheus.Counter
	settingTotal  prometheus.Counter
}

func (s *statsReceiver) HitInc() {
	s.hitCount.Inc()
}

func (s *statsReceiver) ReqInc() {
	s.requestsTotal.Inc()
}

func (s *statsReceiver) ErrorInc() {
	s.errorsTotal.Inc()
}

func (s *statsReceiver) SettingInc() {
	s.settingTotal.Inc()
}
