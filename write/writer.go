package write

import (
	"github.com/joeke80215/dumpcat/pkg"
)

// Writer interface
type Writer interface {
	Write(pkg.Packet) error
}
