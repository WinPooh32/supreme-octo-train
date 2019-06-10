package main

import (
	"image/color"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// lacks := [...]string{
	// " +  К 03.09.16 -",
	// 	" -  П 08.01.16 +  К 25.07.16 -",
	// 	" +  К 22.05.18 -  П 31.05.18 +  К 11.07.18 -  П 03.09.18 +",
	// 	" +  13.07.18 пер(Товар НЕДОСТАЧА В УПАКОВКАХ) -  П 19.07.18 +",
	// 	" + ",
	// 	" +  13.07.18 пер(Товар НЕДОСТАЧА В УПАКОВКАХ) -  П 19.07.18 +",
	// 	" +  22.05.18 пер(Товар НЕДОСТАЧА В УПАКОВКАХ) -  П 11.07.18 +  К 08.10.18 -  П 05.12.18 +",
	// 	" +  К 30.07.18 -",
	// 	" +  01.01.18 пер(ХИМИЯ переделка) -  П 20.02.18 +  К 16.05.18 -  П 15.06.18 +  К 14.08.18 -  П 07.09.18 +",
	// 	" +  30.03.18 пер(ХИМИЯ переделка) -",
	// 	" -  П 29.10.18 + ",
	// 	" Новый товар ",
	// }

	// for _, v := range lacks {
	// 	fmt.Printf("'%s': %v\n", v, parseLackRange(v))
	// }

	serveui("./frontend/dist")
}

// func render(history []float64, item string) {
// 	p, err := plot.New()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Draw a grid behind the data
// 	p.Add(plotter.NewGrid())

// 	p.Title.Text = "Forecast"
// 	p.X.Label.Text = "X"
// 	p.Y.Label.Text = "Y"

// 	length := 52 * 4

// 	// Make a line plotter and set its style.
// 	historyPts := makePlot(history) //randomPoints(length, 50, 0, 0.5, 15, 30)
// 	l, err := plotter.NewLine(historyPts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	l.LineStyle.Width = vg.Points(1)
// 	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

// 	// Make a line plotter and set its style.
// 	smoothHistory := movingavg(history, 2)
// 	upperLimit := confidenceUpperLimit(smoothHistory, 4)

// 	approx := approximateByRegression(upperLimit)
// 	coefs := calcYearCoefficient(history, approx)
// 	multCoefficient(approx, coefs)

// 	coefs2 := calcYearCoefficient(history, approx)
// 	fmt.Println(coefs2)

// 	upperLimitPts := makePlot(approx)
// 	upperLimitLine, upperScatter, err := plotter.NewLinePoints(upperLimitPts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	upperLimitLine.LineStyle.Width = vg.Points(1)
// 	upperLimitLine.LineStyle.Color = color.RGBA{R: 255, B: 255, A: 255}
// 	upperScatter.Shape = draw.PyramidGlyph{}

// 	//smoothPts := makePlot(smoothHistory)
// 	smoothPts := makePlot(upperLimit)
// 	smoothLine, err := plotter.NewLine(smoothPts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	smoothLine.LineStyle.Width = vg.Points(1)
// 	smoothLine.LineStyle.Color = color.RGBA{R: 255, A: 255}

// 	forecastDist := 1

// 	// Make a line plotter and set its style.
// 	// thenTime := time.Now()
// 	// l1, l2 := holtFindParameters(historyPts, forecastDist)
// 	// holtModel, _, _ := buildForecastModelHolt(historyPts, forecastDist, l1, l2)
// 	// nowTime := time.Now()

// 	// forecastLine, err := plotter.NewLine(holtModel)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// forecastLine.LineStyle.Width = vg.Points(1)
// 	// forecastLine.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
// 	// forecastLine.LineStyle.Color = color.RGBA{G: 255, A: 255}

// 	//Trend line
// 	trend := linearRegression(historyPts)
// 	trendPts := make(plotter.XYs, 2)

// 	trendPts[0].X = 0
// 	trendPts[0].Y = trend.Y(0)

// 	trendPts[1].X = float64(len(historyPts))
// 	trendPts[1].Y = trend.Y(int(trendPts[1].X))

// 	fmt.Println(trendPts)

// 	trendLine, err := plotter.NewLine(trendPts)
// 	if err != nil {
// 		panic(err)
// 	}

// 	nextYearLimit := buildForecast(approximateByRegression(upperLimit[0 : len(approx)-52])) //buildForecast(approx[0:len(approx)-52])
// 	nextPlot := makePlot(nextYearLimit)
// 	shiftPlotByX(nextPlot, 52*3)
// 	nextYearLimitLine, err := plotter.NewLine(nextPlot)
// 	if err != nil {
// 		panic(err)
// 	}
// 	nextYearLimitLine.LineStyle.Color = color.RGBA{R: 255, G: 153, A: 255}

// 	p.Add(l, trendLine, smoothLine, upperLimitLine, upperScatter, nextYearLimitLine) //forecastLine

// 	// Set the axis ranges.
// 	p.X.Min = 0
// 	p.X.Max = float64(length + forecastDist)
// 	p.Y.Min = 0
// 	p.Y.Max = findMax(history) + 100

// 	// Save the plot to a PNG file.
// 	if err := p.Save(10*vg.Inch*4, 6*vg.Inch, fmt.Sprintf("item-%s.png", item)); err != nil {
// 		panic(err)
// 	}

// 	// //--------------------------------------
// 	// //Делаем историю кратную 52 неделям
// 	// const layout = "02.01.2006" // dd.mm.yyyy

// 	// beginHistory := "05.03.2015"
// 	// endHistory := "05.09.2019"

// 	// tBegin, _ := time.Parse(layout, beginHistory)
// 	// tEnd, _ := time.Parse(layout, endHistory)

// 	// years := (int(tEnd.Sub(tBegin).Hours())/24)/365
// 	// fmt.Println(years)

// 	// //--
// 	// curWeek := mapWeek("25.05.2019", layout) //dayToWeek(time.Now().YearDay()) // текущая неделя в году
// 	// lastWeeks := (52 - curWeek) // кол-во недель до конца года

// 	// currentHistory := history[curWeek : len(history) - lastWeeks]//История с неполным текущим годом

// 	// simCargo(2000, currentHistory)

// 	fmt.Println("Успех!", item)
// }

func findMax(data []float64) float64 {
	max := 0.0
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return max
}

//Разбиваем на отрезки из трендов
func approximateByRegression(data []float64) []float64 {
	out := make([]float64, len(data))

	weeks := 12
	size := len(data)

	for i := 0; i < size; i++ {
		var period []float64

		if i >= size-2 {
			//регрессия для 2 и менее точек не работает
			out[i] = data[i]
			continue
		} else if i >= size-weeks {
			period = data[i : size-1]
		} else {
			period = data[i : i+weeks-1]
		}

		r := linearRegressionArray(period)
		out[i] = r.Y(0)
	}

	return out
}

func multCoefficient(data []float64, k []float64) {
	years := len(data) / 52

	for i := 0; i < years; i++ {
		coef := k[i]
		for j := 0; j < 52; j++ {
			idx := i*52 + j
			data[idx] *= coef
		}
	}
}

func calcYearCoefficient(dataL, dataR []float64) []float64 {
	sumL := calcYearSums(dataL)
	sumR := calcYearSums(dataR)
	out := make([]float64, len(sumL))

	for i := 0; i < len(sumR); i++ {
		if sumL[i] != 0 && sumR[i] != 0 {
			out[i] = sumL[i] / sumR[i]
		} else {
			out[i] = 0
		}
	}

	return out
}

func calcYearSums(data []float64) []float64 {
	years := len(data) / 52
	sums := make([]float64, years)

	for i := 0; i < years; i++ {
		sum := 0.0
		for j := 0; j < 52; j++ {
			idx := i*52 + j
			sum += data[idx]
		}
		sums[i] = sum
	}

	return sums
}

func simCargo(items int, history []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	p.Title.Text = "Cargo"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Make the cargo plot
	cargoPts := makePlot(simWeeks(52*1, 23, items, buildForecast(history)))
	l, err := plotter.NewLine(cargoPts)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)

	// Set the axis ranges.
	p.X.Min = 0
	p.X.Max = 52
	p.Y.Min = 0
	p.Y.Max = 1000

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch*4, 6*vg.Inch, "cargo.png"); err != nil {
		panic(err)
	}
}

func simWeeks(begin, weeks, items int, forecast []float64) (cargo []float64) {
	cargo = make([]float64, 52)

	for i := begin; i < begin+weeks; i++ {
		items -= int(forecast[i-begin])
		cargo[i-begin] = float64(items)
	}

	return
}

func genNoise(scale, trending float64) float64 {
	return scale * (rand.Float64() - trending)
}

func reverseLacks(lacks []YearLacks) []YearLacks {
	out := make([]YearLacks, len(lacks))
	last := len(out) - 1

	for i, v := range lacks {
		out[last-i] = v
	}

	return out
}

//Реверс только по годам!
func reverse(pts []float64) []float64 {
	out := make([]float64, len(pts))
	// last := len(out) - 1

	years := len(pts) / weeksperyear

	for i := 0; i < years; i++ {
		idx := i * weeksperyear

		b := len(pts) - (i+1)*weeksperyear
		e := b + weeksperyear

		copy(out[idx:idx+weeksperyear], pts[b:e])
	}

	// for i, v := range pts {
	// 	out[last-i] = v
	// }

	return out
}

func shiftPlotByX(plot plotter.XYs, offset int) {
	for i := range plot {
		plot[i].X += float64(offset)
	}
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

// func predictMany(data []float64) []float64 {
// 	const(
// 		l1 = 0.35
// 		l2 = 0.6
// 	)

// 	count := len(data)
// 	out := make([]float64, count)

// 	var a, b float64
// 	// готовим коэффициенты a и b
// 	for i := 0; i < count; i++{
// 		_, a, b = holtForecast(i+1, data[i], a, b ,l1, l2)
// 	}

// 	// экстраполируем
// 	for i := 0; i < count; i++{
// 		next, _, _ := holtForecast(i+1, value, a, b ,l1, l2)
// 		out[i] = next
// 	}

// 	return out
// }

func buildForecast(upperLimit []float64) []float64 {
	years := len(upperLimit) / weeksperyear
	pts := make([]float64, weeksperyear)

	for i := 0; i < weeksperyear; i++ {
		period := zipByYearWeek(upperLimit, i, years, weeksperyear)
		plotPeriod := makePlot(period)

		_, a, b := buildForecastModelHolt(plotPeriod, 1, l1, l2)

		forecast, _, _ := holtForecast(years+1, period[len(period)-1], a, b, l1, l2)

		pts[i] = forecast

		if pts[i] < 0 {
			pts[i] = 0
		}
	}

	return pts
}

func zipByYearWeek(upperLimit []float64, week, years, weeksperyear int) []float64 {
	period := make([]float64, years)
	for i := 0; i < years; i++ {
		period[i] = upperLimit[i*weeksperyear+week]
	}
	return period
}
