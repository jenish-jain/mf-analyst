package analyser

import (
	"fmt"
	"github.com/jenish-jain/mf-analyst/pkg/mfapi"
	"github.com/jenish-jain/mf-analyst/pkg/util"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"math"
	"time"
)

func GetSipPlotChart(sipAmount float64, startDate string, mfData mfapi.MFData) []chart.Series {

	coordinate := GetSipBreakup(sipAmount, startDate, mfData.Data[0].Date, mfData.Data)
	var plotSeries []chart.Series

	for i := 0; i < len(coordinate); i++ {
		var graph chart.TimeSeries

		if i == 0 {
			graph.Name = "Invested amount"
			graph.Style = chart.Style{
				StrokeColor: drawing.ColorGreen,
				FillColor:   drawing.ColorGreen.WithAlpha(35),
			}
			graph.XValues = coordinate[i].Dates
			graph.YValues = coordinate[i].Amounts
		} else {
			graph.Name = "Return Amount"
			graph.Style = chart.Style{
				StrokeColor: drawing.ColorBlue,
				FillColor:   drawing.ColorBlue.WithAlpha(35),
			}
			graph.XValues = coordinate[i].Dates
			graph.YValues = coordinate[i].Amounts
		}
		plotSeries = append(plotSeries, graph)
	}

	return plotSeries
}

func GetSipBreakup(sipAmount float64, sipStartDateString string, sipEndDateString string, historyData []mfapi.Data) []Coordinates {
	var coordinateArray []Coordinates
	var returnCoordinates Coordinates
	startDate := util.GetDateFromDateString("02-01-2006", sipStartDateString)
	endDate := util.GetDateFromDateString("02-01-2006", sipEndDateString)
	//nextInvestmentDate := util.AddMonthsToDate(startDate, 1)
	//var investmentForMonth float64
	//TODO: the current plot is wrong it will hold true only for one investment graph
	// we need to create new series at every investmentCycle and make amount = currentCycleCurrentValuation + prevCycleCurrentValuation

	var currentNav, startNav float64
	startNav = 0
	rawInvestmentCoordinates, investmentMap := getInvestedAmountCoordinates(startDate, endDate, sipAmount, historyData)
	coordinateArray = append(coordinateArray, rawInvestmentCoordinates)

	for i := len(historyData) - 1; i > 0; i-- {
		parsedTime := util.GetDateFromDateString("02-01-2006", historyData[i].Date)
		if parsedTime.After(startDate) && parsedTime.Before(endDate) {
			if startNav == 0 {
				startNav = util.GetFloat64FromString(historyData[i].NAV)
			}
			fmt.Printf("current NAV %f   prev NAV %f InvestedAmount %f time %s \n", startNav, currentNav, investmentMap[parsedTime], historyData[i].Date)
			absoluteReturnRate := (currentNav - startNav) / startNav
			//currentValuation := investmentMap[parsedTime]*(math.Pow(1+absoluteReturnRate,365/util.GetDaysBetweenDates(startDate, parsedTime) ))
			currentValuation := sipAmount * (math.Pow(1+absoluteReturnRate, 365/util.GetDaysBetweenDates(startDate, parsedTime)))
			if investmentMap[parsedTime] != 0 {
				returnCoordinates.Amounts = append(returnCoordinates.Amounts, currentValuation)
				returnCoordinates.Dates = append(returnCoordinates.Dates, parsedTime)
				currentNav = util.GetFloat64FromString(historyData[i].NAV)
			}
		} else {
			currentNav = util.GetFloat64FromString(historyData[i].NAV)
		}
	}
	return append(coordinateArray, returnCoordinates)
}

func getInvestedAmountCoordinates(startDate time.Time, endDate time.Time, sipAmount float64, historyData []mfapi.Data) (Coordinates, map[time.Time]float64) {
	var rawInvestmentCoordinates Coordinates
	var investedAmount float64
	investedAmount = sipAmount
	nextInvestmentDate := util.AddMonthsToDate(startDate, 1)
	investmentMap := make(map[time.Time]float64)

	for i := len(historyData) - 1; i >= 0; i-- {
		mfDate := util.GetDateFromDateString("02-01-2006", historyData[i].Date)
		if mfDate.After(startDate) {
			if mfDate.Before(nextInvestmentDate) {
				rawInvestmentCoordinates.Dates = append(rawInvestmentCoordinates.Dates, mfDate)
				rawInvestmentCoordinates.Amounts = append(rawInvestmentCoordinates.Amounts, investedAmount)
				investmentMap[mfDate] = investedAmount
			} else {
				nextInvestmentDate = util.AddMonthsToDate(mfDate, 1)
				investedAmount = investedAmount + sipAmount
			}
		} else if mfDate.Before(endDate) {
			rawInvestmentCoordinates.Dates = append(rawInvestmentCoordinates.Dates, startDate)
			rawInvestmentCoordinates.Amounts = append(rawInvestmentCoordinates.Amounts, 0)
			investmentMap[startDate] = 0
		} else {
			rawInvestmentCoordinates.Dates = append(rawInvestmentCoordinates.Dates, endDate)
			rawInvestmentCoordinates.Amounts = append(rawInvestmentCoordinates.Amounts, 0)
			investmentMap[endDate] = 0
		}

	}
	return rawInvestmentCoordinates, investmentMap
}

type Coordinates struct {
	Dates   []time.Time `json:"dates"`
	Amounts []float64   `json:"amount"`
}
