package main

import (
	"math"
	"math/big"
)

func bigpowerneg(num *big.Float, exp int) *big.Float {
	if exp == 0 {
		return big.NewFloat(1)
	}

	var accum big.Float
	accum.Copy(big.NewFloat(1))

	for i := 0; i < exp*(-1); i++ {
		var tmp big.Float
		tmp.Copy(&accum)
		accum.Quo(&tmp, num)
	}

	return &accum
}

func bigpower(num *big.Float, exp int) *big.Float {
	if exp == 0 {
		return big.NewFloat(1)
	}

	var accum big.Float
	accum.Copy(num)

	for i := 0; i < exp-1; i++ {
		var tmp big.Float
		tmp.Copy(&accum)
		accum.Mul(&tmp, num)
	}

	return &accum
}

func puasson(value int, avg float64) float64 {
	max := value
	mean := big.NewFloat(float64(avg))

	precalc := bigpowerneg(big.NewFloat(math.E), -int(avg))

	var sum big.Float
	for i := 0; i < max; i++ {
		var tmp big.Int
		var div big.Float

		pow := bigpower(mean, i)
		fact := new(big.Float).SetInt(tmp.MulRange(1, int64(i)))

		div.Quo(pow, fact)
		sum.Add(&sum, &div)
	}

	var tmp big.Float
	tmp.Mul(&sum, precalc)

	ret, _ := sum.Float64()
	return ret
}
