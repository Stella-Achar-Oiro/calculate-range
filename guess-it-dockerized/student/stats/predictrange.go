package stats

// PredictRange predicts the next number range based on the mean and standard deviation of the last N numbers.
func PredictRange(numbers []float64) (float64, float64) {
	mean := Mean(numbers)
	stdDev := StandardDeviation(numbers)
	// Use a prediction range of mean ± 2 * standard deviation
	return mean - 2*stdDev, mean + 2*stdDev
}
