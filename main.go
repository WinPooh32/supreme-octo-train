package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	p.Title.Text = "Forecast"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	length := 300

	// Make a line plotter and set its style.
	historyPts := makePlot(readPoints("продажи.csv")) //randomPoints(length, 50, 0, 0.5, 15, 30)
	l, err := plotter.NewLine(historyPts)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	forecastDist := 1

	// Make a line plotter and set its style.
	thenTime := time.Now()
	l1, l2 := holtFindParameters(historyPts, forecastDist)
	holtModel := buildForecastModelHolt(historyPts, forecastDist, l1, l2)
	nowTime := time.Now()

	forecastLine, err := plotter.NewLine(holtModel)
	if err != nil {
		panic(err)
	}
	forecastLine.LineStyle.Width = vg.Points(1)
	forecastLine.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	forecastLine.LineStyle.Color = color.RGBA{G: 255, A: 255}

	//Trend line
	lineA, lineB := linearRegression(historyPts)
	trendPts := make(plotter.XYs, 2)

	trendPts[0].X = 0
	trendPts[0].Y = lineB

	trendPts[1].X = float64(len(historyPts))
	trendPts[1].Y = lineA*trendPts[1].X + lineB

	fmt.Println(trendPts)

	trendLine, err := plotter.NewLine(trendPts)
	if err != nil {
		panic(err)
	}

	p.Add(l, forecastLine, trendLine)

	// Set the axis ranges.
	p.X.Min = 0
	p.X.Max = float64(length + forecastDist)
	p.Y.Min = 0
	p.Y.Max = 400

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

	fmt.Println("Длина истории:", len(historyPts))
	fmt.Println("Прогноз занял времени:", float64(nowTime.Sub(thenTime).Nanoseconds())*0.000001, "сек.")
	fmt.Println("Средняя ошибка:", calcAvgError(historyPts, holtModel), "%")
	fmt.Println("Успех!")
}

func genNoise(scale, trending float64) float64 {
	return scale * (rand.Float64() - trending)
}

func makePlot(pts []float64) plotter.XYs {
	plotPts := make(plotter.XYs, len(pts))
	for i, v := range pts {
		plotPts[i].X = float64(i)
		plotPts[i].Y = v
	}
	return plotPts
}

// randomPoints returns some random x, y points.
func randomPoints(n, xOffset int, noiseScale, trending, periodStep, scale float64) plotter.XYs {
	pts := make(plotter.XYs, n)

	offset := float64(xOffset)
	trendLast := float64(offset)

	for i := range pts {
		pts[i].X = float64(i)

		trend := trendLast + genNoise(10, trending)
		trendLast = trend

		sin := math.Sin(float64(i) / periodStep)
		// noise := genNoise(noiseScale, trending)

		pts[i].Y = offset + sin*scale + trend
	}

	return pts
}
