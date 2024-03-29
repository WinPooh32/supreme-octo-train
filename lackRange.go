package main

import (
	"regexp"
)

type yearSet [weeksperyear]byte

// LackRange -
type LackRange [2]int

// YearLacks - массивы лет, в которых находятся промежутки провалов
type YearLacks []LackRange

func parseLackRange(s string) YearLacks {
	ranges := make(YearLacks, 0, 5)

	lackInBegin := regexp.MustCompile(`(?s)^\s*\-`)
	lackInEnd := regexp.MustCompile(`(?s).*\-\s*$`)
	dateReg := regexp.MustCompile(`(?s)(\d{2}).(\d{2}).(\d{2})`)
	rangesReg := regexp.MustCompile(`(?s)(\+[^\+]*)|(^\s*\-[^+]*)`)

	rawRanges := rangesReg.FindAllString(s, -1)

	for i, match := range rawRanges {
		var rng LackRange

		dates := dateReg.FindAllString(match, -1)
		size := len(dates)

		if size == 0 {
			// товар есть на конец года
			continue
		}

		isLast := (i == len(rawRanges)-1)

		if size == 1 {
			if lackInBegin.FindStringIndex(s) != nil && lackInEnd.FindStringIndex(match) == nil {
				// нет товара в начале года
				rng[0] = 0
				rng[1] = mapWeek(dates[0], DateLayoutShort)
			} else {
				// нет товара на конец года
				rng[0] = mapWeek(dates[0], DateLayoutShort)
				rng[1] = weeksperyear
			}
		} else if !isLast {
			// нет товара за период
			rng[0] = mapWeek(dates[0], DateLayoutShort)
			rng[1] = mapWeek(dates[1], DateLayoutShort)
		}

		ranges = append(ranges, rng)
	}

	return ranges
}

func fillYearByRanges(lacks YearLacks, out []byte) {
	if len(out) != weeksperyear {
		return
	}

	for _, v := range lacks {
		begin := v[0]
		end := v[1]

		for i := begin; i < end; i++ {
			out[i] = 1
		}
	}
}

func fillManyYearsByLackRanges(yearslacks []YearLacks) []byte {
	out := make([]byte, len(yearslacks)*weeksperyear)
	for i := range yearslacks {
		fillYearByRanges(yearslacks[i], out[i*weeksperyear:(i+1)*weeksperyear])
	}
	return out
}

func restoreLacks(history []float64, lacks []YearLacks) []float64 {
	if len(lacks) > len(history)/weeksperyear {
		panic("кол-во лет в провалах больше, чем всего лет")
	}

	years := len(lacks)

	restored := make([]float64, len(history))
	lacksMapped := fillManyYearsByLackRanges(lacks)

	for i := 0; i < weeksperyear; i++ {

		lack := make([]byte, years)
		pts := make([]float64, 0, years)

		for j := 0; j < years; j++ {
			// индекс = год + неделя
			idx := (j * weeksperyear) + i

			//Заполняем массив точек для построения регрессии
			if lacksMapped[idx] != 1 {
				pts = append(pts, history[idx])
			}
			lack[j] = lacksMapped[idx]
		}

		// восстанавливаем провалы
		l := linearRegressionArray(pts)
		for j, v := range lack {
			if v == 1 {
				pt := l.Y(j)
				if pt > 0 {
					// индекс = год + неделя
					idx := (j * weeksperyear) + i

					restored[idx] = pt

					if history[idx] > 0 && pt > history[idx] {
						history[idx] = (history[idx] + pt) / 2
					} else {
						history[idx] = pt
					}
				}
			}
		}
	}

	return restored
}
