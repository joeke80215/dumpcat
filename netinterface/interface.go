package netinterface

import "github.com/google/gopacket/pcap"

// NetInterface network interface
type NetInterface interface {
	// GetName get network interface name
	GetName() string

	// GetSnapshotLen get network snapshot lenght
	GetSnapshotLen() int32

	// GetPromiscuousMode promiscuous mode
	GetPromiscuousMode() bool
}

type netInterface struct {
	name   string
	sl     int
	pm     bool
	interf pcap.Interface
}

// NewNetInterface new network interface
// parameter:
// name: device name
// sl: snapshot lenght
// pm: promiscuous mode
func NewNetInterface(name string,sl int, pm bool) NetInterface {
	return &netInterface{
		name: name,
		sl: sl,
		pm: pm,
	}
}

// GetName get network interface name
func (n *netInterface) GetName() string {
	return n.name
}

// GetSnapshotLen get network snapshot lenght
func (n *netInterface) GetSnapshotLen() int32 {
	return int32(n.sl)
}

// GetPromiscuousMode promiscuous mode
func (n *netInterface) GetPromiscuousMode() bool {
	return n.pm
}
