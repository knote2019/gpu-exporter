package main

import (
	"fmt"
	exporter "gpu-exporter/pkg/server"
)

func main() {
	fmt.Print("GPU Exporter started !!!\n")
	exporter.StartServer()
	fmt.Print("GPU Exporter started !!!\n")
}
