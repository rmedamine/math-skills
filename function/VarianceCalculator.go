package function

import "math"

func VarianceCalculator(n []float64) float64 {
	nb := AverageCalculator(n)
	result := 0.0
	for _, x := range n {
		result += ((float64(x) - nb) * (float64(x) - nb))
	}
	return math.Round(result / float64(len(n)))
}
