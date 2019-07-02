package main

import "gonum.org/v1/plot/plotter"

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

func makePlot(pts []float64) plotter.XYs {
	plotPts := make(plotter.XYs, len(pts))
	for i, v := range pts {
		plotPts[i].X = float64(i)
		plotPts[i].Y = v
	}
	return plotPts
}

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
