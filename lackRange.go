package main

import (
	"regexp"
)

// LackRange -
type LackRange [2]int

//leftIsLack = regexp.MustCompile(`(?s)^\s*\-`)
//date =  regexp.MustCompile(`(?s)(\d{2}).(\d{2}).(\d{2})`)
//begRange = regexp.MustCompile(`(?s)\-\s*(К|П)?\s*(\d{2}).(\d{2}).(\d{2})\s*\+`)
//endRanges = regexp.MustCompile(`(?s)\+\s*(К|П)?\s*(\d{2}).(\d{2}).(\d{2})\s*\-`)

func parseLackRange(s string) []LackRange {
	ranges := make([]LackRange, 0, 5)

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

		if size == 1 {
			// нет товара в начале года
			rng[0] = 0
			rng[1] = mapWeek(dates[0], DateLayoutShort)
		} else if i != len(rawRanges)-1 {
			// нет товара за период
			rng[0] = mapWeek(dates[0], DateLayoutShort)
			rng[1] = mapWeek(dates[1], DateLayoutShort)
		} else {
			// нет товара на конец года
			rng[0] = mapWeek(dates[0], DateLayoutShort)
			rng[1] = 52
		}

		ranges = append(ranges, rng)
	}

	return ranges
}
