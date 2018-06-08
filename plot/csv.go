package plot

import (
	"fmt"

	"github.com/ksang/automata/rrul"
)

// GenCSV prints RRUL result to CSV format
func GenCSV(result rrul.Result, scale uint) error {
	fmt.Println("Generating CSV result. To plot graph, specify output filename.")
	genTable(result.TCPDownload, "TCPDownload", scale)
	genTable(result.TCPUpload, "TCPUpload", scale)
	genTable(result.UDPRR, "UDPRR", scale)
	return nil
}

func genTable(data []rrul.DataPoint, name string, scale uint) {
	num := len(data)
	fmt.Printf("#%s\n", name)
	for i, dp := range data {
		fmt.Printf("%f,%f\n", float64(scale)*float64(i)/float64(num), dp.Value)
	}
}
