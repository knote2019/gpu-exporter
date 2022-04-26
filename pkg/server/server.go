package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)


func StartServer()  {
	register := prometheus.NewRegistry()
	nums := getGPUNums()
	for i := 0; i < nums; i++ {
		n := strconv.Itoa(i)
		name := getGPUName(n)
		register.MustRegister(NewGPUInfoCollector(n, name))
	}

	http.Handle("/metrics", promhttp.HandlerFor(register, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":12022", nil))
}
