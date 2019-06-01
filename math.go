package main

import (
	"log"
	"math"

	"gonum.org/v1/plot/plotter"
)

type line struct {
	k, b float64
}

func (l *line) Y(x int) float64 {
	return l.k*float64(x) + l.b
}

// Линейная регрессия
// y = kx + b
func linearRegression(history plotter.XYs) line {
	var (
		sumXSquare, sumX, sumY, sumXY float64
	)

	for _, v := range history {
		sumXSquare += v.X * v.X
		sumX += v.X
		sumY += v.Y
		sumXY += v.X * v.Y
	}

	size := float64(len(history))

	// Определитель матрицы системы уравнений
	det := sumXSquare*size - sumX*sumX

	detA := sumXY*size - sumY*sumX
	detB := sumXSquare*sumY - sumX*sumXY

	var l line
	l.k = detA / det
	l.b = detB / det

	return l
}

func linearRegressionArray(history []float64) line {
	var (
		sumXSquare, sumX, sumY, sumXY float64
	)

	for x := 0; x < len(history); x++  {
		sumXSquare += float64(x * x)
		sumX += float64(x)
		sumY += history[x]
		sumXY += float64(x) * history[x]
	}

	size := float64(len(history))

	// Определитель матрицы системы уравнений
	det := sumXSquare*size - sumX*sumX

	detA := sumXY*size - sumY*sumX
	detB := sumXSquare*sumY - sumX*sumXY

	if detA == 0 || det == 0 || detB == 0 {
		return line{}
	}

	var l line
	l.k = detA / det
	l.b = detB / det

	return l
}

// Квадратичное отклонение
func dispersion(data []float64, avg float64) (out []float64) {
	out = make([]float64, len(data))

	for i := range data {
		tmp := data[i] - avg
		out[i] = tmp * tmp
	}

	return
}

// Среднеквадратичное отклонение
func stdDeviation(dispersion []float64) float64 {
	return math.Sqrt(average(dispersion...))
}

// Среднее
func average(nums ...float64) float64 {
	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

// Скользящее среднее
func movingavg(in []float64, window int) (out []float64) {
	if len(in) == 0 {
		log.Fatalln("movingavg(in []float64, window int): пустой массив in!")
	}

	size := len(in)
	out = make([]float64, size)

	for i := 0; i < size; i++ {
		if i < window {
			out[i] = in[i]
			continue
		}

		sum := 0.0
		for j := 0; j < window; j++ {
			sum += in[i-j]
		}

		if sum > 0 {
			out[i] = sum / float64(window)
		}else{
			out[i] = 0
		}
		
	}

	return
}

// func movingavg(in []float64, window int) (out []float64) {
// 	if len(in) == 0 {
// 		log.Fatalln("movingavg(in []float64, window int): пустой массив in!")
// 	}

// 	if window%2 == 0 || window < 3 {
// 		log.Fatalln("movingavg(in []float64, window int): значение window должно быть нечетным числом ")
// 	}

// 	size := len(in)
// 	out = make([]float64, size)

// 	half := window / 2
// 	end := size - half

// 	for i := 0; i < end; i++ {
// 		if i < half {
// 			out[i] = in[i]
// 			continue
// 		}

// 		sum := 0.0

// 		for j := i - half; j < i+half; j++ {
// 			sum += in[j]
// 		}

// 		out[i] = sum / float64(window)
// 	}

// 	return
// }

// Верхняя доверительная граница
func confidenceUpperLimit(in []float64, window int) (out []float64) {
	out = make([]float64, len(in))

	for i := 0; i < len(in); i++ {
		if i < window {
			out[i] = in[i]
			continue
		}

		slice := in[i-window : i]
		deviation := stdDeviation(dispersion(slice, average(slice...)))

		out[i] = in[i] + deviation
	}

	return
}
