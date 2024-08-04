package stats

import "math"

// StandardDeviation calculates the standard deviation from the given variance.
func StandardDeviation(data []float64) float64 {
	mean := Mean(data)
	var sum float64
	for _, value := range data {
		sum += (value - mean) * (value - mean)
	}
	return math.Sqrt(sum / float64(len(data)))
}
