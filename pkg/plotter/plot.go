package plotter

import (
	"fmt"
	"github.com/jenish-jain/mf-analyst/pkg/mfapi"
	"github.com/wcharczuk/go-chart/v2"
	"log"
	"os"
	"time"
)

func PlotTimeSeries(Title string, XValues []time.Time, YValues []float64, data mfapi.MFData) {
	priceSeries := chart.TimeSeries{
		Name: Title,
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: XValues,
		YValues: YValues,
	}
	//smaSeries := chart.SMASeries{
	//	Name: "SPY - SMA",
	//	Style: chart.Style{
	//		StrokeColor:     drawing.ColorRed,
	//		StrokeDashArray: []float64{5.0, 5.0},
	//	},
	//	InnerSeries: priceSeries,
	//}
	//
	//bbSeries := &chart.BollingerBandsSeries{
	//	Name: "SPY - Bol. Bands",
	//	Style: chart.Style{
	//		StrokeColor: drawing.ColorFromHex("efefef"),
	//		FillColor:   drawing.ColorFromHex("efefef").WithAlpha(64),
	//	},
	//	InnerSeries: priceSeries,
	//}

	//sipSeries := analyser.GetSipPlotChart(5000, data)

	graph := chart.Chart{
		Title:      Title,
		TitleStyle: chart.Shown(),
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionUnderTick,
		},
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Max: 100,
				Min: 0.0,
			},
		},
		Series: []chart.Series{
			//bbSeries,
			priceSeries,
			//&sipSeries,
			//smaSeries,
			//chart.AnnotationSeries{
			//	Annotations: []chart.Value2{
			//		{XValue: 1.0, YValue: 5.0},
			//	},
			//},
		},
		Width:  2000,
		Height: 800,
	}

	err := os.Mkdir("charts", 0755)
	if err != nil {
		log.Printf("error making new directory - %v", err)
	}
	outFileName := fmt.Sprintf("./charts/%s.png", Title)
	f, err := os.Create(outFileName)
	if err != nil {
		log.Printf("error creating file - %v", err)
	}
	defer f.Close()
	err = graph.Render(chart.PNG, f)
	if err!= nil {
		log.Printf("error rendering graph - %v", err)
	}
}
