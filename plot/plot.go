/*
Package plot provides plot facilities to visualize RRUL data
*/
package plot

import "github.com/ksang/automata/rrul"

// Config is the configuration of plot package
type Config struct {
	Filename string
	Scale    uint
}

// Visualize is the entry point of plot package
func Visualize(cfg Config, data rrul.Result) error {
	defer Summary(data)
	if cfg.Filename == "" {
		return GenCSV(data, cfg.Scale)
	}
	return LineChart(data, cfg.Scale, cfg.Filename)
}
