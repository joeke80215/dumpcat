package main

import (
	"runtime"

	"github.com/joeke80215/dumpcat/exec"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	exec.Exec()
}
