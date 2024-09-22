package metrics

import (
	"log"
	"sync"

	"github.com/mindprince/gonvml"
)

var (
	initOnce        sync.Once
	initializeError error
	deviceCount     uint
)

func GetGPUMetrics() ([]float64, []float64) {
	// init NVML
	initOnce.Do(func() {
		initializeError = gonvml.Initialize()
		if initializeError != nil {
			log.Printf("Failed to initialize NVML: %v", initializeError)
			return
		}
		var err error
		deviceCount, err = gonvml.DeviceCount()
		if err != nil {
			log.Printf("Failed to get device count: %v", err)
			initializeError = err
			return
		}
	})

	if initializeError != nil {
		return nil, nil
	}

	// get metrics of each device
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
