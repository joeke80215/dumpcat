package translate

import (
	"log"

	"github.com/google/gopacket"
	"github.com/joeke80215/dumpcat/networklayer"
	"github.com/joeke80215/dumpcat/pkg"
)

// Translate translate package
type Translate interface {
	// Tran translate gopacket.Packet to pkg.Packet
	Tran(pkg.Packet) pkg.Packet
}

type translate struct {
	layer    networklayer.NetworkLayer
	dumpName string
}

// NewTranslate new package translate
// proto "ipv4" "tcp"
func NewTranslate(dumpName string, proto string, packet gopacket.Packet) Translate {
	t := &translate{
		dumpName: dumpName,
	}
	switch proto {
	case "ipv4":
		t.layer = networklayer.NewIPv4Layer(packet)
	case "tcp":
		t.layer = networklayer.NewtcpLayer(packet)
	default:
		log.Fatal("translate protocal error proto=", proto)
	}

	return t
}

// Tran translate gopacket.Packet to pkg.PKG
func (t *translate) Tran(p pkg.Packet) pkg.Packet {
	if !t.layer.CheckLayer() || !t.layer.CheckPackageLayer() {
		return p
	}
	p.Set("DumpName", t.dumpName)
	return t.layer.Tran(p)
}
