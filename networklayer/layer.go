package networklayer

import (
	"github.com/google/gopacket"
	"github.com/joeke80215/dumpcat/pkg"
)

// NetworkLayer network layer
type NetworkLayer interface {
	// Tran translate gopacket.Packet to pkg.Packet
	Tran(pkg.Packet) pkg.Packet

	// CheckLayer check layer is exsist,if exsist return true or false
	CheckLayer() bool

	// CheckPackageLayer check package layer is exsist,if exsist return true or false
	CheckPackageLayer() bool

	// SupportLayer support next high level layer
	// SupportLayer() []string
}

type layer struct {
	packet gopacket.Packet
}
