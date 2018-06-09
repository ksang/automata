package plot

import (
	"fmt"

	"github.com/ksang/automata/rrul"
)

// Summary prints out sumary result of RRUL test
func Summary(result rrul.Result) {
	fmt.Printf("++Summary:\n")
	fmt.Printf("  TCP Upload Throughput:\t%f Mbps\n", result.TCPUploadThroughput)
	fmt.Printf("  TCP Download Throughput:\t%f Mbps\n", result.TCPDownloadThroughput)
	fmt.Printf("  UDP Round-Robin Latency:\t%f ms\n", result.UDPRRThroughput)
}
