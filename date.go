package main

import (
	"log"
	"math"
	"time"
)

//
const (
	DateLayoutLong  = "02.01.2006" // dd.mm.yyyy
	DateLayoutShort = "02.01.06"   // dd.mm.yy
)

func dayToWeek(day int) (week int) {
	week = int(math.Floor(float64(day) / 7.0))
	return
}

func mapWeek(date string, layout string) (week int) {

	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalln(err)
	}

	week = dayToWeek(t.YearDay())
	return
}
