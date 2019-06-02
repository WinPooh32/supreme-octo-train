package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func toFloat64(in []string) []float64 {
	out := make([]float64, len(in))

	for i := range in {
		num, err := strconv.ParseFloat(in[i], 64)
		if err != nil {
			// log.Fatalln(err)
			num = 0
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

func readPoints(filename string, line int) []float64 {
	fileIn, _ := os.Open(filename)
	defer fileIn.Close()
	reader := csv.NewReader(fileIn)

	for i := 0; ;i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i == line {
			return toFloat64(record)
		}
	}

	return nil
}