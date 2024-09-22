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

	netRX = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "net_rx_bytes",
			Help: "Net interface received in byte",
		},
		[]string{"interface"},
	)

	netTX = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "net_tx_bytes",
			Help: "Net interface transmited in byte",
		},
		[]string{"interface"},
	)
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memUsage)
	prometheus.MustRegister(gpuUsage)
	prometheus.MustRegister(gpuMemUsage)
	prometheus.MustRegister(netRX)
	prometheus.MustRegister(netTX)
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

	//** Network **//
	rx, _ := metrics.GetNetworkRXBytes()
	for i := 0; i < len(rx); i++ {
		netRX.With(prometheus.Labels{"interface": rx[i].Name}).Set(float64(rx[i].Bytes))
	}

	tx, _ := metrics.GetNetworkTXBytes()
	for i := 0; i < len(rx); i++ {
		netTX.With(prometheus.Labels{"interface": tx[i].Name}).Set(float64(tx[i].Bytes))
	}
}
