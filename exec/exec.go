package exec

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/joeke80215/dumpcat/config"
	"github.com/joeke80215/dumpcat/dump"
	"github.com/joeke80215/dumpcat/filter"
	"github.com/joeke80215/dumpcat/handle"
	"github.com/joeke80215/dumpcat/netinterface"
	"github.com/joeke80215/dumpcat/pkg"
	"github.com/joeke80215/dumpcat/task"
	"github.com/joeke80215/dumpcat/translate"
)

var (
	block chan bool
)

func init() {
	block = make(chan bool)
}

// Exec execute dump package to log
func Exec() {
	// show libcap version
	fmt.Println(pcap.Version())
	// show network interface list
	var devices []pcap.Interface
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range devices {
		fmt.Println(v)
	}

	tasker := task.NewTask()
	tasker.Register(config.Cfg.GetOutput())

	for index := range config.Cfg.GetDumpList() {
		net := netinterface.NewNetInterface(config.Cfg.GetDeviceName(index), 65535, false)
		h := handle.NewHandler(net)
		h.SetBPFFilter(filter.NewFilter(config.Cfg.GetBPF(index)))
		packets := dump.NewDumper().Dump(h)
		ipv4 := pkg.NewPacket()

		go func(name string) {
			for packet := range packets {
				tn := time.Now()
				ipv4.Set("Time", &tn)
				mPKG := translate.NewTranslate(name, "ipv4", packet).Tran(ipv4)
				mPKG = translate.NewTranslate(name, "tcp", packet).Tran(mPKG)
				for _, layer := range tasker.GetLogics() {
					mPKG = layer.Middleware(mPKG)
				}
				for _, w := range tasker.GetWriters() {
					err := w.Write(mPKG)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}(index)
	}

	<-block
}
