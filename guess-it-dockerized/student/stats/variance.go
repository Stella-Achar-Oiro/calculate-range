package stats

// Variance calculates the variance of a slice of float64 numbers.
func Variance(data []float64, mean float64) float64 {
    var sumOfSquares float64
    for _, value := range data {
        sumOfSquares += (value - mean) * (value - mean)
    }
    return sumOfSquares / float64(len(data))
}
