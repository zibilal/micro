package main

import (
	"github.com/mataharimall/micro"
	_ "github.com/mataharimall/micro/apps/loket"
	"github.com/mataharimall/micro/config"
	"log"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := config.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	micro.RouterManager.Init()
}
