package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percentage",
			Help: "CPU usage percentage per core",
		},
		[]string{"core"},
	)

	memUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Memory usage in bytes",
		},
	)

	gpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gpu_usage_percentage",
			Help: "GPU usage in bytes",
		},
		[]string{"GPU"},
	)

	gpuMemUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gpu_memory_usage_bytes",
			Help: "GPU memory usage in bytes",
		},
		[]string{"GPU"},
	)
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memUsage)
	prometheus.MustRegister(gpuUsage)
	prometheus.MustRegister(gpuMemUsage)
}

func getMetrics() {
	//** CPU usage **//
	coreUsage := GetCPUUsage()
	for i := 0; i < len(coreUsage); i++ {
		label := fmt.Sprintf("%d", i)
		cpuUsage.With(prometheus.Labels{"core": label}).Set(coreUsage[i])
	}

	//** main memory usage **//
	memUsage.Set(GetMemoryUsage())

	gpuUtil, gpuMemUsed := GetGPUMetrics()
	for i := 0; i < len(gpuUtil); i++ {
		label := fmt.Sprintf("%d", i)
		gpuUsage.With(prometheus.Labels{"GPU": label}).Set(gpuUtil[i])
		gpuMemUsage.With(prometheus.Labels{"GPU": label}).Set(gpuMemUsed[i])
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/metrics")

	getMetrics()
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	log.Println("HPC node exporter runs on http://localhost:8080/metrics")
	http.HandleFunc("/metrics", metricsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
