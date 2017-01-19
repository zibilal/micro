package main

import (
	"fmt"
	"github.com/mataharimall/micro/app"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	if err := app.Construct(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
