package analyser

import (
	"github.com/jenish-jain/mf-analyst/pkg/mfapi"
	"github.com/wcharczuk/go-chart/v2"
	"strconv"
	"time"
)

func getSipPlotData(sipAmount float64, mfData mfapi.MFData) ([]time.Time, []float64) {

	historyData := mfData.Data

	var dates []time.Time
	for _, ts := range historyData {
		parsed, _ := time.Parse("02-01-2006", ts.Date)
		dates = append(dates, parsed)
	}

	var values []float64
	var sipValues []float64

	for i := 0; i < len(historyData); i++ {
		rawValue, _ := strconv.ParseFloat(historyData[i].NAV, 64)
		values = append(values, rawValue)
		if i > 0 {
			bucketInterestRate := values[i] / values[i-1]
			sipValue := sipAmount * bucketInterestRate
			sipValues = append(sipValues, sipValue)
		} else {
			sipValues = append(sipValues, 0)
		}

	}
	//fmt.Println(sipValues)
	//fmt.Println(len(sipValues))
	//fmt.Println(len(dates))

	return dates, sipValues
}

func GetSipPlotChart(sipAmount float64, mfData mfapi.MFData) chart.TimeSeries {

	xValues, yValues := getSipPlotData(sipAmount, mfData)

	graph := chart.TimeSeries{
		Name: mfData.Meta.SchemeName,
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xValues,
		YValues: yValues,
	}

	return graph
}
