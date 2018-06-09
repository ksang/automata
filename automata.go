package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ksang/automata/plot"
	"github.com/ksang/automata/rrul"
)

var (
	// netserver address
	host string
	// netserver base port
	port uint
	// time in seconds to test
	last uint
	// output file name of plotting
	output string
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "netserver address e.g \"192.168.100.100\"")
	flag.UintVar(&port, "p", 12865, "netserver base port, port+1 and port+2 are also used")
	flag.UintVar(&last, "l", 30, "time in seconds the test will last")
	flag.StringVar(&output, "o", "", "output filename of plotting, print csv data if not provided")
}

func checkPort(p uint) bool {
	return p <= 65535
}

func main() {
	flag.Parse()
	if !checkPort(port) {
		fmt.Fprintln(os.Stderr, "Invalid port:", port)
		flag.PrintDefaults()
		return
	}
	res, err := rrul.Launch(
		rrul.Config{
			Host:    host,
			Port:    port,
			Seconds: last,
		})
	if err != nil {
		log.Fatal(err)
	}
	if err := plot.Visualize(plot.Config{
		Filename: output,
		Scale:    last,
	}, res); err != nil {
		log.Fatal(err)
	}
}
