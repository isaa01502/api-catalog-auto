package metrics

import "github.com/prometheus/client_golang/prometheus"

// Init Инициализирует метрики
func Init() {
	prometheus.MustRegister(getOrCreateRedisRequestDuration())
}
