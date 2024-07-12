package stats

import (
	"math"
	"sort"
)

// Mean calculates the average of a slice of float64 numbers.
func Mean(data []float64) float64 {
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

// Median calculates the median value of a slice of float64 numbers.
func Median(data []float64) float64 {
	// Create a copy of the input data to avoid modifying the original slice.
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	// Sort the copied slice in ascending order.
	sort.Float64s(sortedData)
	n := len(sortedData)
	if n%2 == 0 {
		midIndex := n / 2
		return (sortedData[midIndex-1] + sortedData[midIndex]) / 2.0
	}
	return sortedData[n/2]
}

// Variance calculates the variance of a slice of float64 numbers given the mean.
func Variance(data []float64, mean float64) float64 {
	var totalVar float64
	for _, num := range data {
		diff := num - mean
		totalVar += diff * diff
	}
	return totalVar / float64(len(data))
}

// StandardDeviation calculates the standard deviation from the given variance.
func StandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
