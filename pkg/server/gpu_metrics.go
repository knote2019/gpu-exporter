package server

import (
	"fmt"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type GPUInfoCollector struct {
	Seq            string
	Temperature    *prometheus.Desc
	MemTotal       *prometheus.Desc
	MemUsed        *prometheus.Desc
	MemFree        *prometheus.Desc
	GPUUtilization *prometheus.Desc
}

func NewGPUInfoCollector(num string, name string) *GPUInfoCollector {
	return &GPUInfoCollector{
		Seq: num,
		Temperature: prometheus.NewDesc(
			"gpu_temperature",
			"Shows temperature about gpu",
			nil,
			prometheus.Labels{"gpu_seq": num, "name": name}),
		MemTotal: prometheus.NewDesc(
			"gpu_memory_total",
			"Shows gpu memory total (MiB)",
			nil,
			prometheus.Labels{"gpu_seq": num, "name": name}),
		MemUsed: prometheus.NewDesc(
			"gpu_memory_used",
			"Shows gpu memory used (MiB)",
			nil,
			prometheus.Labels{"gpu_seq": num, "name": name}),
		MemFree: prometheus.NewDesc(
			"gpu_memory_free",
			"Shows gpu memory free (MiB)",
			nil,
			prometheus.Labels{"gpu_seq": num, "name": name}),
		GPUUtilization: prometheus.NewDesc(
			"gpu_utilization",
			"Shows gpu utilization (%)",
			nil,
			prometheus.Labels{"gpu_seq": num, "name": name}),
	}
}

func (c *GPUInfoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Temperature
	ch <- c.MemTotal
	ch <- c.MemUsed
	ch <- c.MemFree
	ch <- c.GPUUtilization
}

func (c *GPUInfoCollector) Collect(ch chan<- prometheus.Metric) {

	tempVal := getGPUTemperature(c.Seq)
	totalVal := getGPUMemTotal(c.Seq)
	usedVal := getGPUMemUsed(c.Seq)
	freeVal := getGPUMemFree(c.Seq)
	utilizationVal := getGPUUtilization(c.Seq)

	ch <- prometheus.MustNewConstMetric(c.Temperature, prometheus.GaugeValue, tempVal)
	ch <- prometheus.MustNewConstMetric(c.MemTotal, prometheus.GaugeValue, totalVal)
	ch <- prometheus.MustNewConstMetric(c.MemUsed, prometheus.GaugeValue, usedVal)
	ch <- prometheus.MustNewConstMetric(c.MemFree, prometheus.GaugeValue, freeVal)
	ch <- prometheus.MustNewConstMetric(c.GPUUtilization, prometheus.GaugeValue, utilizationVal)
}

// create NVML handle.
func createNvmlHandle(){
	ret := nvml.Init()
	if ret != nvml.SUCCESS{
		fmt.Print("create nvml handle failed")
	}
}

// delete NVML handle.
func deleteNvmlHandle(){
	ret := nvml.Shutdown()
	if ret != nvml.SUCCESS{
		fmt.Print("delete nvml handle failed")
	}
}

// get gpu temperature.
func getGPUTemperature(n string) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	index,_ := strconv.Atoi(n)
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuTemperature,_ := nvml.DeviceGetTemperature(device, 0)
	return float64(gpuTemperature)
}

// get GPU mem total (MiB)
func getGPUMemTotal(n string) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	index,_ := strconv.Atoi(n)
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuMemoryInfo,_ := nvml.DeviceGetMemoryInfo(device)
	return float64(gpuMemoryInfo.Total / 1024 / 1024)
}

// get GPU mem used (MiB)
func getGPUMemUsed(n string) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	index,_ := strconv.Atoi(n)
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuMemoryInfo,_ := nvml.DeviceGetMemoryInfo(device)
	return float64(gpuMemoryInfo.Used / 1024 / 1024)
}

// get GPU mem free (MiB)
func getGPUMemFree(n string) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	index,_ := strconv.Atoi(n)
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuMemoryInfo,_ := nvml.DeviceGetMemoryInfo(device)
	return float64(gpuMemoryInfo.Free / 1024 / 1024)
}

// get GPU utilization.
func getGPUUtilization(n string) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	index,_ := strconv.Atoi(n)
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuUtilizeInfo,_ := nvml.DeviceGetUtilizationRates(device)
	return float64(gpuUtilizeInfo.Gpu)
}

// get GPU name.
func getGPUName(n string) string {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	index,_ := strconv.Atoi(n)
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuName,_ := nvml.DeviceGetName(device)
	return gpuName
}

// get GPU number.
func getGPUNums() int {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	count, _ := nvml.DeviceGetCount()
	return count
}
