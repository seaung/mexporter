package main

import (
	"log"
	"net/http"

	"mexport/collector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metrics = "/metrics"
)

func main() {
	mexport := collector.NewMexportCollector("")
	prometheus.MustRegister(mexport)
	http.Handle(metrics, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			`<html>
             <head><title>Mexporter</title></head>
             <body>
             <h1>Mirth Channel Exporter</h1>
             <p><a href='` + metrics + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	log.Fatal(http.ListenAndServe(":9527", nil))
}
