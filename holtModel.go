package main

import (
	"log"
	"math"

	"gonum.org/v1/plot/plotter"
)

//Модель Хольта
func holtForecast(nextX int, currentValue, a, b, l1, l2 float64) (forecast, newA, newB float64) {
	if nextX < 1 {
		log.Fatal()
	}

	forecast = a + b*float64(nextX)
	newA = l1*currentValue + (1.0-l1)*(a-b)
	newB = l2*(newA-a) + (1.0-l2)*b

	return
}

func calcAvgError(history, forecast plotter.XYs) float64 {
	forecastDist := int(forecast[0].X - history[0].X)
	sum := 0.0

	for i := forecastDist; i < len(history); i++ {
		real := history[i].Y

		if real < 1.0 {
			continue
		}

		predicted := forecast[i-forecastDist].Y

		sum += 100.0 * math.Abs(real-predicted) / real
	}

	return sum / float64(len(history)-forecastDist)
}

func buildForecastModelHolt(history plotter.XYs, forecastDist int, l1, l2 float64) plotter.XYs {
	size := len(history)
	modelPts := make(plotter.XYs, size)

	a := history[0].Y
	b := 0.0

	for i := 0; i < size; i++ {
		old := history[i]
		next, newA, newB := holtForecast(forecastDist, old.Y, a, b, l1, l2)

		a = newA
		b = newB

		modelPts[i].X = float64(i + forecastDist)
		modelPts[i].Y = next
	}

	return modelPts
}

//Поиск оптимальных параметров
func holtFindParameters(history plotter.XYs, forecastDist int) (minL1, minL2 float64) {
	minErr := math.MaxFloat64
	minL1 = 0.0
	minL2 = 0.0

	const step = 0.1

	// т.к. l1 и l2 зависят друг от друга, то нужно перебрать все варианты
	// и лишь потом взять константы, дающие минимальную ошибку.
	for l1 := 0.0; l1 < 1.0; l1 += step {
		for l2 := 0.0; l2 < 1.0; l2 += step {
			forecast := buildForecastModelHolt(history, forecastDist, l1, l2)
			err := calcAvgError(history, forecast)

			if err < minErr {
				minErr = err
				minL1 = l1
				minL2 = l2
			}
		}
	}

	return
}
