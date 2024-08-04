package stats

import "sort"

// Median calculates the median value of a slice of float64 numbers.
func Median(data []float64) float64 {
	n := len(data)
	sortedData := make([]float64, n)
	copy(sortedData, data)
	sort.Float64s(sortedData)

	if n%2 == 0 {
		return (sortedData[n/2-1] + sortedData[n/2]) / 2
	}
	return sortedData[n/2]
}
