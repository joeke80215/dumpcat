package task

import (
	"github.com/joeke80215/dumpcat/config"
	"github.com/joeke80215/dumpcat/elasticsearch"
	"github.com/joeke80215/dumpcat/logic"
	"github.com/joeke80215/dumpcat/write"
)

// Task write and logic layer task
type Task interface {
	// Register register writer and logic layer
	Register(map[string]config.Output) Task

	// GetLogics get logic layers
	GetWriters() []write.Writer

	// GetWriters get writers
	GetLogics() []logic.Logic
}

type task struct {
	writers []write.Writer
	logics  []logic.Logic
}

// NewTask new task write and logic layer task
func NewTask() Task {
	return &task{}
}

// Register register writer and logic layer
func (t *task) Register(svcs map[string]config.Output) Task {
	for i := range svcs {
		switch i {
		case "elasticsearch":
			t.writers = append(t.writers, elasticsearch.NewESConn())
		}
	}

	for _, v := range config.Cfg.GetLogics() {
		switch v {
		case "latency":
			t.logics = append(t.logics, logic.NewLatency())
		}
	}

	return t
}

// GetLogics get logic layers
func (t *task) GetLogics() []logic.Logic {
	return t.logics
}

// GetWriters get writers
func (t *task) GetWriters() []write.Writer {
	return t.writers
}
