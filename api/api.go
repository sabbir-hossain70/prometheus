package api

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "api_ping_count",
		Help: "Number of ping requests",
	})

func ping(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "sabbir")
}

func RunServer(Port string) {
	fmt.Println("Inside RunServer")
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+Port, nil)
}
