package ui

import (
	"fmt"
	stat2 "github.com/vit1251/golden/pkg/stat"
	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"net/http"
)

type StatImageAction struct {
	Action
}

func NewStatImageAction() *StatImageAction {
	sa := new(StatImageAction)
	return sa
}

func (self *StatImageAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var statManager *stat2.StatManager
	self.Container.Invoke(func(sm *stat2.StatManager) {
		statManager = sm
	})

	/* Get statistics */
	sums, err1 := statManager.GetMessageSummary()
	if err1 != nil {
		panic(err1)
	}

	/* Chart values */
	var chartValues []chart.Value

	for _, v := range sums {
		cv1 := chart.Value{
			Value: float64(v.Value),
			Label: fmt.Sprintf("%s", v.Date),
		}
		chartValues = append(chartValues, cv1)
	}

	/* Render chart */
	graph := chart.BarChart{
		Title: "RX message statistics",
		TitleStyle: chart.Style{
			Show:                true,
		},
		Height: 480,
		Width: 640,
		BarWidth: 25,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    25,
				Left:   25,
				Right:  25,
				Bottom: 50,
			},
			FillColor: drawing.ColorFromHex("C0C0C0"),
		},
		XAxis: chart.Style{
			Show:                true,
			TextRotationDegrees: 90.0,
		},
		YAxis: chart.YAxis{
			Name: "RX message count",
			NameStyle: chart.Style{
				Show:                true,
			},
			Style: chart.Style{
				Show:                true,
			},
		},
		Bars: chartValues,
	}

	w.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, w)

}
