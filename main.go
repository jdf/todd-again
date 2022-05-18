package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/jdf/todd-again/todd"
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

	todd.Run()
}
