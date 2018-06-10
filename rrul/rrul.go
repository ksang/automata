/*
Package rrul implements RRUL(Real-Time Response Under Load) networking performance test
based on netperf utilities
*/
package rrul

// DataPoint defines a time-series data point
type DataPoint struct {
	Time  float64
	Value float64
	Unit  string
}

// Result defines time-series data returned by RRUL test
type Result struct {
	TCPUpload             []DataPoint
	TCPDownload           []DataPoint
	UDPRR                 []DataPoint
	ICMPRR                []DataPoint
	TCPUploadThroughput   float64
	TCPDownloadThroughput float64
	UDPRRThroughput       float64
	ICMPRRThroughput      float64
}

// Launch is the entry point of running RRUL test
func Launch(conf Config) (Result, error) {
	n := &netperf{
		cfg:    conf,
		data:   Result{},
		quitCh: make(chan struct{}, 1),
	}
	if err := n.start(); err != nil {
		return Result{}, err
	}
	return n.data, nil
}
