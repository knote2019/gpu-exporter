package main

import (
	"fmt"
	gpu_monitor "gpu-monitor/pkg/server"
)

func main() {
	fmt.Print("GPU Monitor started !!!\n")
	gpu_monitor.StartServer()
	fmt.Print("GPU Monitor started !!!\n")
}
