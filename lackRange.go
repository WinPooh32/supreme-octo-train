package main

import (
	"regexp"
)

// LackRange -
type LackRange [2]int

func parseLackRange(s string) []LackRange {
	ranges := make([]LackRange, 0, 5)

	lackInBegin := regexp.MustCompile(`(?s)^\s*\-`)
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
			if lackInBegin.FindStringIndex(s) != nil {
				// нет товара в начале года
				rng[0] = 0
				rng[1] = mapWeek(dates[0], DateLayoutShort)
			}else{
				// нет товара на конец года
				rng[0] = mapWeek(dates[0], DateLayoutShort)
				rng[1] = 52
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
