package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

var redisRequestDuration *prometheus.HistogramVec
var redisRequestDurationMutex sync.Mutex

func RedisRequestDuration(start time.Time, action, method string, err error) {
	if redisRequestDuration == nil {
		return
	}
	elapsed := float64(time.Since(start)) / float64(time.Second)
	status := "ok"
	if err != nil {
		status = "error"
	}
	redisRequestDuration.WithLabelValues(action, method, status).Observe(elapsed)
}

func getOrCreateRedisRequestDuration() *prometheus.HistogramVec {
	if redisRequestDuration == nil {
		redisRequestDurationMutex.Lock()
		defer redisRequestDurationMutex.Unlock()
		if redisRequestDuration != nil {
			return redisRequestDuration
		}
		redisRequestDuration = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: "hboauth",
				Name:      "redis_request_duration",
				Help:      "Redis request durations, status",
				Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 0.75, 1, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 5.0, 10.0},
			}, []string{"action", "method", "status"})
	}
	return redisRequestDuration
}
