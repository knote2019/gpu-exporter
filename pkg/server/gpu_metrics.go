package server

import (
	"fmt"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type GPUInfoCollector struct {
	Seq            int
	Temperature    *prometheus.Desc
	MemTotal       *prometheus.Desc
	MemUsed        *prometheus.Desc
	MemFree        *prometheus.Desc
	MemUtilization *prometheus.Desc
	GPUUtilization *prometheus.Desc
}

func NewGPUInfoCollector(index int, name string, uuid string) *GPUInfoCollector {
	return &GPUInfoCollector{
		Seq: index,
		Temperature: prometheus.NewDesc(
			"gpu_temperature",
			"Shows gpu temperature (C)",
			nil,
			prometheus.Labels{"gpu_seq": strconv.Itoa(index), "name": name, "uuid": uuid}),
		MemTotal: prometheus.NewDesc(
			"gpu_memory_total",
			"Shows gpu memory total (MiB)",
			nil,
			prometheus.Labels{"gpu_seq": strconv.Itoa(index), "name": name, "uuid": uuid}),
		MemUsed: prometheus.NewDesc(
			"gpu_memory_used",
			"Shows gpu memory used (MiB)",
			nil,
			prometheus.Labels{"gpu_seq": strconv.Itoa(index), "name": name, "uuid": uuid}),
		MemFree: prometheus.NewDesc(
			"gpu_memory_free",
			"Shows gpu memory free (MiB)",
			nil,
			prometheus.Labels{"gpu_seq": strconv.Itoa(index), "name": name, "uuid": uuid}),
		MemUtilization: prometheus.NewDesc(
			"mem_utilization",
			"Shows mem utilization (%)",
			nil,
			prometheus.Labels{"gpu_seq": strconv.Itoa(index), "name": name, "uuid": uuid}),
		GPUUtilization: prometheus.NewDesc(
			"gpu_utilization",
			"Shows gpu utilization (%)",
			nil,
			prometheus.Labels{"gpu_seq": strconv.Itoa(index), "name": name, "uuid": uuid}),
	}
}

func (c *GPUInfoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Temperature
	ch <- c.MemTotal
	ch <- c.MemUsed
	ch <- c.MemFree
	ch <- c.MemUtilization
	ch <- c.GPUUtilization
}

func (c *GPUInfoCollector) Collect(ch chan<- prometheus.Metric) {

	tempVal := getGPUTemperature(c.Seq)
	totalVal := getGPUMemTotal(c.Seq)
	usedVal := getGPUMemUsed(c.Seq)
	freeVal := getGPUMemFree(c.Seq)
	MemUtilizationVal := getMemUtilization(c.Seq)
	GpuUtilizationVal := getGPUUtilization(c.Seq)

	ch <- prometheus.MustNewConstMetric(c.Temperature, prometheus.GaugeValue, tempVal)
	ch <- prometheus.MustNewConstMetric(c.MemTotal, prometheus.GaugeValue, totalVal)
	ch <- prometheus.MustNewConstMetric(c.MemUsed, prometheus.GaugeValue, usedVal)
	ch <- prometheus.MustNewConstMetric(c.MemFree, prometheus.GaugeValue, freeVal)
	ch <- prometheus.MustNewConstMetric(c.MemUtilization, prometheus.GaugeValue, MemUtilizationVal)
	ch <- prometheus.MustNewConstMetric(c.GPUUtilization, prometheus.GaugeValue, GpuUtilizationVal)
}

// create NVML handle.
func createNvmlHandle(){
	//fmt.Printf("nvml.Init() ...")
	ret := nvml.Init()
	if ret != nvml.SUCCESS{
		fmt.Print("create nvml handle failed")
	}
}

// delete NVML handle.
func deleteNvmlHandle(){
	//fmt.Printf("nvml.Shutdown() ...")
	ret := nvml.Shutdown()
	if ret != nvml.SUCCESS{
		fmt.Print("delete nvml handle failed")
	}
}


// get GPU number.
func getGPUNums() int {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	count, _ := nvml.DeviceGetCount()
	return count
}

// get GPU name.
func getGPUName(index int) string {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuName,_ := nvml.DeviceGetName(device)
	return gpuName
}

// get GPU name.
func getGPUUuid(index int) string {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuUuid,_ := nvml.DeviceGetUUID(device)
	return gpuUuid
}

// get gpu temperature.
func getGPUTemperature(index int) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuTemperature,_ := nvml.DeviceGetTemperature(device, 0)
	return float64(gpuTemperature)
}

// get GPU mem total (MiB)
func getGPUMemTotal(index int) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuMemoryInfo,_ := nvml.DeviceGetMemoryInfo(device)
	return float64(gpuMemoryInfo.Total / 1024 / 1024)
}

// get GPU mem used (MiB)
func getGPUMemUsed(index int) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuMemoryInfo,_ := nvml.DeviceGetMemoryInfo(device)
	return float64(gpuMemoryInfo.Used / 1024 / 1024)
}

// get GPU mem free (MiB)
func getGPUMemFree(index int) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuMemoryInfo,_ := nvml.DeviceGetMemoryInfo(device)
	return float64(gpuMemoryInfo.Free / 1024 / 1024)
}

// get Mem utilization.
func getMemUtilization(index int) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuUtilizeInfo,_ := nvml.DeviceGetUtilizationRates(device)
	return float64(gpuUtilizeInfo.Memory)
}

// get GPU utilization.
func getGPUUtilization(index int) float64 {
	createNvmlHandle()
	defer deleteNvmlHandle()
	// GET START.
	device,_ := nvml.DeviceGetHandleByIndex(index)
	gpuUtilizeInfo,_ := nvml.DeviceGetUtilizationRates(device)
	return float64(gpuUtilizeInfo.Gpu)
}
