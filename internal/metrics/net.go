package metrics

import (
	"github.com/shirou/gopsutil/net"
)

type InterfaceUsage struct {
	Name  string
	Bytes uint64
}

func GetNetworkRXBytes() ([]InterfaceUsage, error) {
	interfaces, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	var txData []InterfaceUsage
	for _, iface := range interfaces {
		txData = append(txData, InterfaceUsage{
			Name:  iface.Name,
			Bytes: iface.BytesRecv,
		})
	}

	return txData, nil
}

func GetNetworkTXBytes() ([]InterfaceUsage, error) {
	interfaces, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	var txData []InterfaceUsage
	for _, iface := range interfaces {
		txData = append(txData, InterfaceUsage{
			Name:  iface.Name,
			Bytes: iface.BytesSent,
		})
	}

	return txData, nil
}
