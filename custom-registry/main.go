package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a non-global / non-default registry.
	registry := prometheus.NewRegistry()
	// OPTIONAL: Add process and Go runtime metrics to our custom registry.
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())

	// Create a single gauge without any labels.
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	// Register the gauge with our metrics registry.
	registry.MustRegister(temp)

	// Set the gauge's value to 42.
	temp.Set(42)

	// Expose our custom registry over HTTP on /metrics.
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
