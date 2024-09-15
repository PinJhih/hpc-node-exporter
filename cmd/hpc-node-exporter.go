package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"hpc-node-exporter/internal/exporter"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/metrics")
	exporter.GetMetrics()
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	log.Println("HPC node exporter runs on http://0.0.0.0:8080/metrics")
	http.HandleFunc("/metrics", metricsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
