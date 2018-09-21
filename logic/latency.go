package logic

import (
	"sync"
	"time"

	"github.com/joeke80215/dumpcat/pkg"
)

/*
caculate connection latest packet latency
*/

type latency struct {
	mut  sync.Mutex
	pkgs map[string]*state
}

type state struct {
	src  string
	dst  string
	time *time.Time
}

// NewLatency new time offset logic strage
func NewLatency() Logic {
	return &latency{
		pkgs: make(map[string]*state),
	}
}

// Middleware get last packet time and now time offset
// add parameter "Latency" (float64)
func (lat *latency) Middleware(pkg pkg.Packet) pkg.Packet {
	lat.mut.Lock()
	defer lat.mut.Unlock()
	tt := pkg.Get("Time").(*time.Time)
	seq0 := pkg.Get("Src").(string) + pkg.Get("Dst").(string)
	seq1 := pkg.Get("Dst").(string) + pkg.Get("Src").(string)
	seq := ""
	if lseq0 := lat.pkgs[seq0]; lseq0 != nil {
		seq = seq0
	} else if lseq1 := lat.pkgs[seq1]; lseq1 != nil {
		seq = seq1
	} else {
		lat.pkgs[seq0] = &state{
			src:  pkg.Get("Src").(string),
			dst:  pkg.Get("Dst").(string),
			time: tt,
		}
		pkg.Set("Latency", float64(0))
		return pkg
	}

	if (pkg.Get("SYN").(bool) && !pkg.Get("ACK").(bool)) || pkg.Get("FIN").(bool) {
		pkg.Set("Latency", float64(0))
		lat.pkgs[seq0] = &state{
			src:  pkg.Get("Src").(string),
			dst:  pkg.Get("Dst").(string),
			time: tt,
		}
	} else {
		if lat.pkgs[seq].dst == pkg.Get("Src").(string) {
			subs := tt.Sub(*lat.pkgs[seq].time)
			pkg.Set("Latency", subs.Seconds())
		} else if lat.pkgs[seq].src == pkg.Get("Src").(string) {
			pkg.Set("Latency", float64(0))
		}
	}

	lat.pkgs[seq].time = tt

	return pkg
}
