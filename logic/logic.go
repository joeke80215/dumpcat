package logic

import (
	"github.com/joeke80215/dumpcat/pkg"
)

var (
	// LogicPacket logic layer channel
	LogicPacket chan pkg.Packet
)

// Logic middleware interface
type Logic interface {
	Middleware(pkg.Packet) pkg.Packet
}

func init() {
	LogicPacket = make(chan pkg.Packet)
}
