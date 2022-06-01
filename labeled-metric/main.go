package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a gauge with two label names ("house" and "room").
	temp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// The two label names by which to split the metric.
		[]string{"house", "room"},
	)

	// Register the gauge with our metrics registry.
	prometheus.MustRegister(temp)

	// Set the temperature to different values, depending on house and room.
	temp.WithLabelValues("julius", "living-room").Set(23.5)
	temp.WithLabelValues("julius", "bedroom").Set(21.3)
	temp.WithLabelValues("fred", "living-room").Set(21.5)
	temp.WithLabelValues("fred", "bedroom").Set(20.7)

	// Expose our custom registry over HTTP on /metrics.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
