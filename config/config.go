package config

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	// Cfg config info
	Cfg   configer
	fPath string
)

type configer interface {
	// GetDumpList get dump list
	GetDumpList() map[string]Dump

	// GetDeviceName get index of list device name
	GetDeviceName(string) string

	// GetBPF get index of list BPF
	GetBPF(string) string

	// GetLogics get logic layer list
	GetLogics() []string

	// GetOutput get output list
	GetOutput() map[string]Output
}

// Config config info
type Config struct {
	DumpList map[string]Dump   `yaml:"dumpList"`
	Logics   []string          `yaml:"logics"`
	Outputs  map[string]Output `yaml:"output"`
}

// Dump dump info
type Dump struct {
	Device string `yaml:"device"`
	BPF    string `yaml:"bpf"`
}

// Output output distination
type Output struct {
	Host     string            `yaml:"host"`
	User     string            `yaml:"user,omitempty"`
	Password string            `yaml:"password,omitempty"`
	Options  map[string]string `yaml:"options,omitempty"`
}

func init() {
	flag.StringVar(&fPath, "f", "./config.yaml", "config.yaml path")
	flag.Parse()
	Cfg = &Config{}

	f, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(f, Cfg)
	if err != nil {
		log.Fatal(err)
	}
}

// GetDumpList get dump list
func (c *Config) GetDumpList() map[string]Dump {
	return c.DumpList
}

// GetDeviceName get index of list device name
func (c *Config) GetDeviceName(i string) string {
	return c.DumpList[i].Device
}

// GetBPF get index of list BPF
func (c *Config) GetBPF(i string) string {
	return c.DumpList[i].BPF
}

// GetLogics get logic layer list
func (c *Config) GetLogics() []string {
	return c.Logics
}

// GetOutput get output list
func (c *Config) GetOutput() map[string]Output {
	return c.Outputs
}
