package dump

import (
	"github.com/google/gopacket"
	"github.com/joeke80215/dumpcat/handle"
)

// Dumper package info
type Dumper interface {
	// Dump and send package
	Dump(handle.Handler) chan gopacket.Packet
}

type dumper struct{}

// NewDumper new dumper
func NewDumper() Dumper {
	return &dumper{}
}

// Dump and send package
func (dp *dumper) Dump(handle handle.Handler) chan gopacket.Packet {
	packetSource := gopacket.NewPacketSource(
		handle.GetHandle(),
		handle.GetHandle().LinkType(),
	)

	return packetSource.Packets()
}
