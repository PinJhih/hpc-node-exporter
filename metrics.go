package main

import (
	"log"

	"github.com/mindprince/gonvml"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func GetCPUUsage() []float64 {
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		log.Printf("Error getting CPU usage: %v", err)
		return nil
	}
	return percentages
}

func GetMemoryUsage() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory usage: %v", err)
		return 0.0
	}
	return float64(v.Used)
}

func GetGPUMetrics() ([]float64, []float64) {
	
	if err := gonvml.Initialize(); err != nil {
		log.Printf("Failed to initialize NVML: %v", err)
		return nil, nil
	}
	defer gonvml.Shutdown()

	deviceCount, err := gonvml.DeviceCount()
	if err != nil {
		log.Printf("Failed to get device count: %v", err)
		return nil, nil
	}

	gpuUsage := make([]float64, deviceCount)
	gpuMemoryUsage := make([]float64, deviceCount)

	for i := uint(0); i < deviceCount; i++ {
		device, err := gonvml.DeviceHandleByIndex(i)
		if err != nil {
			log.Printf("Failed to get device handle for device %d: %v", i, err)
			continue
		}

		utilization, _, err := device.UtilizationRates()
		if err != nil {
			log.Printf("Failed to get utilization rates for device %d: %v", i, err)
			continue
		}
		gpuUsage[i] = float64(utilization)

		_, used, err := device.MemoryInfo()
		if err != nil {
			log.Printf("Failed to get memory info for device %d: %v", i, err)
			continue
		}
		gpuMemoryUsage[i] = float64(used)
	}

	return gpuUsage, gpuMemoryUsage
}
