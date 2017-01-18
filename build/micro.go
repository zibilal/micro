package main

import (
	"github.com/mataharimall/micro"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	micro.Construct()
}
