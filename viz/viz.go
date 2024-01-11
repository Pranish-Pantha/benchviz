package viz

import (
	"os"

	"github.com/Pranish-Pantha/benchviz/benchmark"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Viz interface {
	Render(file string)
}

type LineSeries struct {
	name  string
	items []opts.LineData
}

type BarSeries struct {
	name  string
	items []opts.BarData
}

type LineChart struct {
	title    string
	subtitle string
	xAxis    []string
	series   []LineSeries
}

type BarChart struct {
	title    string
	subtitle string
	xAxis    []string
	series   []BarSeries
}

// NewLineChart creates a new line chart with the given title, subtitle, x-axis labels, and series
func NewLineChart(title, subtitle string, xAxis []string, series []LineSeries) *LineChart {
	return &LineChart{
		title:    title,
		subtitle: subtitle,
		xAxis:    xAxis,
		series:   series,
	}
}

// NewBarChart creates a new bar chart with the given title, subtitle, x-axis labels, and series
func NewBarChart(title, subtitle string, xAxis []string, series []BarSeries) *BarChart {
	return &BarChart{
		title:    title,
		subtitle: subtitle,
		xAxis:    xAxis,
		series:   series,
	}
}

// NewLineSeries creates a new line series with the given name and items
func NewLineSeries(name string, items []opts.LineData) LineSeries {
	return LineSeries{
		name:  name,
		items: items,
	}
}

// NewBarSeries creates a new bar series with the given name and items
func NewBarSeries(name string, items []opts.BarData) BarSeries {
	return BarSeries{
		name:  name,
		items: items,
	}
}

// Render renders the line chart to the given file
func (lc *LineChart) Render(file string) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    lc.title,
			Subtitle: lc.subtitle,
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "# of elements", NameLocation: "center"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Time (ms)", NameLocation: "center"}),
	)
	line.SetXAxis(lc.xAxis)
	for _, s := range lc.series {
		line.AddSeries(s.name, s.items)
	}
	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	line.Render(f)
}

// Render renders the bar chart to the given file
func (bc *BarChart) Render(file string) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    bc.title,
			Subtitle: bc.subtitle,
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "# of elements", NameLocation: "center"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Time (ms)", NameLocation: "center"}),
	)
	bar.SetXAxis(bc.xAxis)
	bar.SetSeriesOptions(charts.WithLabelOpts(opts.Label{Show: true}))
	for _, s := range bc.series {
		bar.AddSeries(s.name, s.items)
	}

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	bar.Render(f)
}

// BenchmarkToLineSeries converts a benchmark to a line series
func BenchmarkToLineSeries(bench benchmark.Benchmark) LineSeries {
	items := make([]opts.LineData, 0)
	for _, testCase := range bench.Tests {
		items = append(items, opts.LineData{Value: testCase.NanosecondsPerOp / 1000000}) // ms/op
	}
	return NewLineSeries(bench.Name, items)
}

// BenchmarkToBarSeries converts a benchmark to a bar series
func BenchmarkToBarSeries(bench benchmark.Benchmark) BarSeries {
	items := make([]opts.BarData, 0)
	for _, testCase := range bench.Tests {
		items = append(items, opts.BarData{Value: testCase.NanosecondsPerOp / 1000000}) // ms/op
	}
	return NewBarSeries(bench.Name, items)
}
