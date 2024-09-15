package metrics

import (
	"log"

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
