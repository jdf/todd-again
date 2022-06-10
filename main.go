package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/jdf/todd-again/engine"
	"github.com/jdf/todd-again/testgame"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	engine.RunGameLoop(testgame.Level1(), 1600, 900, "Todd Again")
}
