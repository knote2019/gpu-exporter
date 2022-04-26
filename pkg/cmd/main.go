package main

import (
	"fmt"
	"gpu-exporter/pkg/exporter"
)

func main() {
	fmt.Print("GPU Exporter started !!!\n")
	exporter.StartServer()
	fmt.Print("GPU Exporter started !!!\n")
}
