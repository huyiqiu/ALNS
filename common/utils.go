package common

import (
	"math/rand"
	"time"
)

// 
func RouletteSelect(roulette map[string]float64) string {
	// // random select
	// i := len(roulette)
	// j := 0
	// var random string
	// i2 := RandInt(1, i)
	// for k := range roulette {
	// 	j ++
	// 	if j == i2 {
	// 		random = k
	// 		break
	// 	}
	// }
	// return random

	// RouletteSelect select
	sum := 0.0
	for _, v := range roulette {
		sum += v
	}
	position := RandDecimal() * sum
	vernier := 0.0
	var random string
	for k, v := range roulette {
		random = k
		vernier += v
		if position <= vernier {
			return k
		}
	}
	return random
}

func RandInt(start, end int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return rng.Intn(end - start + 1) + start
}

func RandDecimal() float64 {
	return float64(RandInt(0, 100)) / 100.0
}

func NotIn(set []int, i int) bool {
	for _, v := range set {
		if v == i {
			return false
		}
	}
	return true
}