package networklayer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/joeke80215/dumpcat/pkg"
)

// ipv4Layer ipv4 packet layer
type ipv4Layer struct {
	layer *layer
}

// NewIPv4Layer new ipv4 layer
func NewIPv4Layer(packet gopacket.Packet) NetworkLayer {
	return &ipv4Layer{
		layer: &layer{
			packet: packet,
		},
	}
}

// CheckLayer check tcp layer is exsist,if exsist return true or false
func (ipv4Layer *ipv4Layer) CheckLayer() bool {
	if ipv4Layer.layer.packet.Layer(layers.LayerTypeIPv4) == nil {
		return false
	}

	return true
}

// CheckPackageLayer check ipv4 package layer is exsist,if exsist return true or false
func (ipv4Layer *ipv4Layer) CheckPackageLayer() bool {
	if err := ipv4Layer.layer.packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4); err == nil {
		return false
	}

	return true
}

// Tran translate ipv4 gopacket.Packet to pkg.PKG
func (ipv4Layer *ipv4Layer) Tran(p pkg.Packet) pkg.Packet {
	ipLayer := ipv4Layer.layer.packet.Layer(layers.LayerTypeIPv4)
	ipv4, _ := ipLayer.(*layers.IPv4)
	p.Set("Src", ipv4.SrcIP.String())
	p.Set("Dst", ipv4.DstIP.String())
	p.Set("Proto", "ipv4")
	p.Set("Length", ipv4.Length)

	return p
}
