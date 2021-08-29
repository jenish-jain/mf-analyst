package main

//go:generate go run main.go

import (
	"github.com/jenish-jain/mf-analyst/pkg/mfapi"
	"github.com/jenish-jain/mf-analyst/pkg/plotter"
	"strconv"
	"time"
)

func main() {
	mfData := mfapi.GetMfData(120468)
	xv, yv := xvalues(mfData.Data), yvalues(mfData.Data)

	plotter.PlotTimeSeries(mfData.Meta.SchemeName, xv, yv, mfData)

}

func xvalues(data []mfapi.Data) []time.Time {

	var dates []time.Time
	for _, ts := range data {
		parsed, _ := time.Parse("02-01-2006", ts.Date)
		dates = append(dates, parsed)
	}
	return dates
}

func yvalues(data []mfapi.Data) []float64 {
	var values []float64
	for _, val := range data {
		parsed, _ := strconv.ParseFloat(val.NAV, 64)
		values = append(values, parsed)
	}
	return values

}
