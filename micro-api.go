package main

import (
	_ "github.com/mataharimall/micro-api/apps/loket"
	"github.com/mataharimall/micro-api/commons"
	"github.com/mataharimall/micro-api/config"
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
	commons.RouterManager.Init()
}
