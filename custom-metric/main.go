package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a gauge metric without any label dimensions.
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	// Register the metric with the default registry.
	prometheus.MustRegister(temp)

	// Set the gauge's value to 42.
	temp.Set(42)

	// Expose the metrics on /metrics on port 8080.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
