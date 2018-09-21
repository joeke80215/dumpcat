package device

import "github.com/google/gopacket/pcap"

// Device network interface controller device
type Device interface {
	// GetList get devices list
	GetList() []pcap.Interface

	// GetDevice get network interface device
	GetDevice(int) pcap.Interface
}

type device struct {
	devices []pcap.Interface
}

// NewDevice new network device
func NewDevice() Device {
	return &device{}
}

// GetList get devices list
func (d *device) GetList() []pcap.Interface {
	return d.devices
}

// GetDevice get network interface device
func (d *device) GetDevice(i int) pcap.Interface {
	return d.devices[i]
}
