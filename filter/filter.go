package filter

import "sync"

// Filter filter package
type Filter interface {
	GetBPF() string
}

// filter data
type filter struct {
	ori   string
	rwMut sync.RWMutex
}

// NewFilter new filter ori is BPF format
// reference: http://biot.com/capstats/bpf.html
func NewFilter(ori string) Filter {
	return &filter{ori: ori}
}

// GetBPF get BPF format string
func (f *filter) GetBPF() string {
	return f.ori
}
