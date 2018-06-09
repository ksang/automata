package plot

import (
	"fmt"
	"image/color"
	"os"

	"github.com/ksang/automata/rrul"
	"github.com/pplcc/plotext"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// LineChart generate line chart for RRUL result and save a PNG file
func LineChart(data rrul.Result, scale uint, filename string) error {

	pTraffic, err := plot.New()
	if err != nil {
		return err
	}
	pTraffic.X.Label.Text = "Time(s)"
	pTraffic.Y.Label.Text = "Mbps"
	pTraffic.Y.Tick.Marker = commaTicks{}

	if err = plotutil.AddLinePoints(pTraffic,
		"TCP Upload", MakePoints(data.TCPUpload, scale),
		"TCP Download", MakePoints(data.TCPDownload, scale)); err != nil {
		return err
	}

	pLatancy, err := plot.New()
	if err != nil {
		return err
	}

	pLatancy.X.Label.Text = "Time(s)"
	pLatancy.Y.Label.Text = "Latency(ms)"

	// Make a line plotter and set its style.
	udpLine, udpPoints, err := plotter.NewLinePoints(MakePoints(data.UDPRR, scale))
	if err != nil {
		return err
	}
	udpLine.Color = color.RGBA{41, 141, 255, 255}
	udpPoints.Color = color.RGBA{41, 141, 255, 255}
	udpPoints.Shape = draw.PyramidGlyph{}
	pLatancy.Add(udpLine, udpPoints)
	pLatancy.Legend.Add("UDP Round-Robin Latency", udpLine, udpPoints)
	// table is used for align two graph
	table := plotext.Table{
		RowHeights: []float64{1, 1},
		ColWidths:  []float64{1},
	}
	plots := [][]*plot.Plot{[]*plot.Plot{pTraffic}, []*plot.Plot{pLatancy}}
	img := vgimg.New(800, 600)
	dc := draw.New(img)

	canvases := table.Align(plots, dc)
	plots[0][0].Draw(canvases[0][0])
	plots[1][0].Draw(canvases[1][0])

	w, err := os.Create(filename)
	if err != nil {
		return err
	}

	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(w); err != nil {
		return err
	}
	fmt.Println("LineChart saved to", filename)
	return nil
}

// MakePoints converts a RRUL datapoint list to plotters points
func MakePoints(data []rrul.DataPoint, scale uint) plotter.XYs {
	n := len(data)
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = float64(scale) * float64(i) / float64(n)
		pts[i].Y = data[i].Value
	}
	return pts
}

type commaTicks struct{}

// Ticks computes the default tick marks, but inserts commas
// into the labels for the major tick marks.
func (commaTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)
	for i, t := range tks {
		if t.Label == "" { // Skip minor ticks, they are fine.
			continue
		}
		tks[i].Label = addCommas(t.Label)
	}
	return tks
}

// AddCommas adds commas after every 3 characters from right to left.
// NOTE: This function is a quick hack, it doesn't work with decimal
// points, and may have a bunch of other problems.
func addCommas(s string) string {
	rev := ""
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
		n++
		if n%3 == 0 {
			rev += ","
		}
	}
	s = ""
	for i := len(rev) - 1; i >= 0; i-- {
		s += string(rev[i])
	}
	return s
}
