package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gonum.org/v1/plot/plotter"
)

//y = ax + b
func linearRegression(history plotter.XYs) (a, b float64) {
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

	a = detA / det
	b = detB / det

	return
}

func toFloat64(in []string) []float64 {
	out := make([]float64, len(in))

	for i := range in {
		num, err := strconv.ParseFloat(in[i], 64)
		if err != nil {
			log.Fatalln(err)
		}

		out[i] = num
	}

	return out
}

func toString(in []float64) []string {
	prt := message.NewPrinter(language.Russian)

	out := make([]string, len(in))

	for i := range in {
		num := prt.Sprintf("%f", in[i])
		out[i] = num
	}
	return out
}

func readPoints(filename string) []float64 {
	fileIn, _ := os.Open(filename)
	defer fileIn.Close()
	reader := csv.NewReader(fileIn)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		return toFloat64(record)
	}

	return nil
}
