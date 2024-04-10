package http

import (
	"git.homebank.kz/homebank-halykid/api-halykid/internal/metrics"
	redisstats "git.homebank.kz/libs/go-prometheus-go-redis.v8-stats"
	promgin "git.homebank.kz/libs/prometheus-gin"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *server) addMetrics() {
	metrics.Init()
	promgin.Use(s.router, "hbhalykid")

	redisProm := redisstats.NewStatsCollector(s.Manager.Cache, "hbhalykid", "", "")
	prometheus.MustRegister(redisProm)
}
