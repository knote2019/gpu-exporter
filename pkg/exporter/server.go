package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)


func StartServer()  {
	register := prometheus.NewRegistry()
	nums := getGPUNums()
	for index := 0; index < nums; index++ {
		name := getGPUName(index)
		uuid := getGPUUuid(index)
		register.MustRegister(NewGPUInfoCollector(index, name, uuid))
	}
	// start exporter.
	http.Handle("/metrics", promhttp.HandlerFor(register, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":12022", nil))
}
