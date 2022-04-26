package main

import (
	"fmt"
	gpu_monitor "gpu-monitor/pkg/server"
)

func main() {
	fmt.Print("GPU Exporter started !!!\n")
	gpu_monitor.StartServer()
	fmt.Print("GPU Exporter started !!!\n")
}
