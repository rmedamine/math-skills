package function

import (
	"math"
	"sort"
)

func MedianCalculator(x []float64) float64 {
	sort.Float64s(x)

	size := len(x)

	if size == 0 {
		return 0
	}
	if size%2 == 0 {
		return float64(x[(size/2)-1]+x[(size/2)]) / 2.0
	}
	return float64(math.Round(float64(x[(size)/2])))
}
