package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/dweb-personal-website/resources/templates"
	"github.com/prometheus/client_golang/prometheus"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "endpoint"},
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	totalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
	component := templates.IndexPage()
	component.Render(context.Background(), w)
}
