package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"os"
)

var (
	env          string
	counterVec   *prometheus.CounterVec
	guageVec     *prometheus.GaugeVec
	summaryVec   *prometheus.SummaryVec
	histogramVec *prometheus.HistogramVec
)

func init() {
	counterVec = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_received",
			Help: "Number of HTTP requests received.",
		},
		[]string{"package", "server", "method", "env"},
	)

	guageVec = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "processing_time",
			Help: "Number of HTTP responses sent.",
		},
		[]string{"package", "server", "method", "env"},
	)

	summaryVec = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "summary_vec",
			Help: "summary_vec_help",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.99: 0.001,
			},
		},
		[]string{"package", "server", "method", "code", "env"},
	)

	histogramVec = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_durations_ms_histogram",
			Help:    "HTTP latency distributions histogram.",
			Buckets: []float64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096},
		},
		[]string{"package", "server", "method", "code", "env"},
	)

	os.Setenv("APP_ENV", "local")
	env = os.Getenv("APP_ENV")

}

func Metrics(packageName, server, method, code string, processingTime int64) {

	IncCounterVec(packageName, server, method, env)

	ObserveHistogramVec(packageName, server, method, code, processingTime)

	ObserveSummaryVec(packageName, server, method, code, processingTime)

	SetGuageVec(packageName, server, method, env, processingTime)
}
