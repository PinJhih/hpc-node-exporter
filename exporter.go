package main

import (
    "fmt"
    "log"
    "net/http"
    "runtime"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
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
)

func init() {
    prometheus.MustRegister(cpuUsage)
    prometheus.MustRegister(memUsage)
}

func getMetrics() {
    for i := 0; i < runtime.NumCPU(); i++ {
        util := 0.0 // TODO: get CPU core utilization
        cpuUsage.With(prometheus.Labels{"core": fmt.Sprintf("%d", i)}).Set(util)
    }

    memUsage.Set(0.0) // TODO: get RAM usage
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
