package function

import "math"

func StandardDeviation(n []float64) float64 {
	nb := VarianceCalculator(n)

	return math.Round(math.Sqrt(nb))
}
