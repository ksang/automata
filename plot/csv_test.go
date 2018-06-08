package plot

import (
	"testing"

	"github.com/ksang/automata/rrul"
)

func TestGenCSV(t *testing.T) {
	tests := []struct {
		res   rrul.Result
		scale uint
	}{
		{
			rrul.Result{},
			0,
		},
		{
			rrul.Result{
				TCPUpload: []rrul.DataPoint{
					rrul.DataPoint{
						Value: 1.1,
					},
					rrul.DataPoint{
						Value: 1.2,
					},
					rrul.DataPoint{
						Value: 1.3,
					},
					rrul.DataPoint{
						Value: 5555,
					},
				},
				TCPDownload: []rrul.DataPoint{
					rrul.DataPoint{
						Value: 31.1,
					},
					rrul.DataPoint{
						Value: 1.2,
					},
					rrul.DataPoint{
						Value: 1.3,
					},
				},
				UDPRR: []rrul.DataPoint{},
			},
			60,
		},
	}
	for _, te := range tests {
		if err := GenCSV(te.res, te.scale); err != nil {
			t.Errorf("error: %v\n", err)
		}
	}
}
