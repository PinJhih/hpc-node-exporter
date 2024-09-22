package exporter

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"hpc-node-exporter/internal/metrics"
)

var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percentage",
			Help: "CPU utilization of each core",
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
			Help: "GPU utilization of each GPU",
		},
		[]string{"GPU"},
	)

	gpuMemUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gpu_memory_usage_bytes",
			Help: "GPU memory usage in bytes of each GPU",
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

func GetMetrics() {
	//** CPU usage **//
	coreUsage := metrics.GetCPUUsage()
	for i := 0; i < len(coreUsage); i++ {
		label := fmt.Sprintf("%d", i)
		cpuUsage.With(prometheus.Labels{"core": label}).Set(coreUsage[i])
	}

	//** main memory usage **//
	memUsage.Set(metrics.GetMemoryUsage())

	
	//** GPU usage **//
	gpuUtil, gpuMemUsed := metrics.GetGPUMetrics()
	for i := 0; i < len(gpuUtil); i++ {
		label := fmt.Sprintf("%d", i)
		gpuUsage.With(prometheus.Labels{"GPU": label}).Set(gpuUtil[i])
		gpuMemUsage.With(prometheus.Labels{"GPU": label}).Set(gpuMemUsed[i])
	}
}
