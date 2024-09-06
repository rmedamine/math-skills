package function

import "math"

func AverageCalculator(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	result := 0.0

	for _, n := range x {
		result += n
	}
	return math.Round(float64(result) / float64(len(x)-1))
}
