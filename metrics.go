package main

import (
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func GetCPUUsage(core int) float64 {
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		log.Printf("Error getting CPU usage: %v", err)
		return 0.0
	}
	if core < len(percentages) {
		return percentages[core]
	}
	return 0.0
}

func GetMemoryUsage() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory usage: %v", err)
		return 0.0
	}
	return float64(v.Used)
}
