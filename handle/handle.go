package handle

import (
	"log"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/joeke80215/dumpcat/filter"
	"github.com/joeke80215/dumpcat/netinterface"
)

// Handler network package handler
type Handler interface {
	// SetBPFFilter set BPF filter
	// reference: http://biot.com/capstats/bpf.html
	SetBPFFilter(filter.Filter)

	// GetHandle get handle
	GetHandle() *pcap.Handle

	// Close handler
	Close()
}

type handler struct {
	handle *pcap.Handle
	err    error
}

// NewHandler new network package handler
// reference: https://blog.lab99.org/post/golang-2017-09-23-video-packet-capture-analysis-with-go.html
func NewHandler(net netinterface.NetInterface) Handler {
	h := &handler{}
	h.handle, h.err = pcap.OpenLive(
		net.GetName(),            // device name
		net.GetSnapshotLen(),     //	snapshot length
		net.GetPromiscuousMode(), //	promiscuous mode?
		-1*time.Second,           // timeout 负数表示不缓存，直接输出
	)
	if h.err != nil {
		log.Fatal(h.err)
	}
	return h
}

// GetHandle get handle
func (h *handler) GetHandle() *pcap.Handle {
	return h.handle
}

// Close handler
func (h *handler) Close() {
	h.handle.Close()
}

// SetBPFFilter set BPF filter
// reference: http://biot.com/capstats/bpf.html
func (h *handler) SetBPFFilter(bpf filter.Filter) {
	if err := h.handle.SetBPFFilter(bpf.GetBPF()); err != nil {
		log.Fatal(err)
	}
}
