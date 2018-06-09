package rrul

import (
	"strconv"
	"strings"
)

// MarshalOutput marshals a command line output from netperf command and turns into []DataPoint
// e.g.
//		Interim result:    0.24 10^6bits/s over 0.199 seconds ending at 1528437383.138
func MarshalOutput(output []byte) ([]DataPoint, error) {
	ret := []DataPoint{}
	rows := strings.Split(string(output), "\n")
	for _, row := range rows {
		dp := DataPoint{}
		idx1 := strings.Index(row, "Interim result:")
		if idx1 < 0 {
			continue
		}
		idx2 := idx1 + 15
		for i := idx2; i < len(row); i++ {
			if row[i] != '\t' && row[i] != ' ' {
				idx2 = i
				break
			}
		}
		idx3 := strings.Index(row[idx2:], " ") + idx2
		dp.Value, _ = strconv.ParseFloat(strings.TrimSpace(row[idx2:idx3]), 10)
		idx4 := strings.Index(row[idx3+1:], " ") + idx3 + 1
		dp.Unit = strings.TrimSpace(row[idx3+1 : idx4])
		idx5 := strings.Index(row[idx4:], "ending at") + idx4 + 9
		dp.Time, _ = strconv.ParseFloat(strings.TrimSpace(row[idx5:]), 10)
		ret = append(ret, dp)
	}
	return ret, nil
}
