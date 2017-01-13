package main

import (
	"github.com/mataharimall/micro-api/config"
	"log"
	"os"
	"runtime"
)

const (
	APP_VER = "v1"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := config.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
}
