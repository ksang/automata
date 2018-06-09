package rrul

import (
	"strconv"
	"strings"
)

// MarshalOutput marshals a command line output from netperf command and turns into []DataPoint
// e.g.
//		NETPERF_INTERIM_RESULT[0]=0.35
//		NETPERF_UNITS[0]=10^6bits/s
//		NETPERF_INTERVAL[0]=0.200
//		NETPERF_ENDING[0]=1528545723.525
func MarshalOutput(output []byte) ([]DataPoint, float64, error) {
	ret := []DataPoint{}
	rows := strings.Split(string(output), "\n")
	var throughput float64
	for i := 0; i < len(rows); i++ {
		if strings.HasPrefix(rows[i], "NETPERF_INTERIM_RESULT") {
			if i+3 >= len(rows) {
				continue
			}
			// value
			idx := strings.Index(rows[i], "=")
			if idx < 0 {
				continue
			}
			value, err := strconv.ParseFloat(rows[i][idx+1:], 10)
			if err != nil {
				continue
			}
			// unit
			idx = strings.Index(rows[i+1], "=")
			unit := rows[i+1][idx+1:]
			// time
			idx = strings.Index(rows[i+3], "=")
			time, _ := strconv.ParseFloat(rows[i+3][idx+1:], 10)
			ret = append(ret, DataPoint{
				Value: value,
				Unit:  unit,
				Time:  time,
			})
			i += 3
		} else if strings.HasPrefix(rows[i], "THROUGHPUT") {
			idx := strings.Index(rows[i], "=")
			if idx < 0 {
				continue
			}
			throughput, _ = strconv.ParseFloat(rows[i][idx+1:], 10)
		}
	}
	return ret, throughput, nil
}
