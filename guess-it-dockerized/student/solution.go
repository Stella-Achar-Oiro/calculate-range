package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"guess-it-1/stats"
)

// Read data from a file specified as a command-line argument.
func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go data.txt")
		return
	}
	var data []float64
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Exiting...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		data = append(data, float64(num))
		if len(data) > 1 {
			lower, upper := calculateRange(data)
			fmt.Printf("%d %d\n", lower, upper)
		}

	}
}

// Function to calculate range for the next number
func calculateRange(data []float64) (int, int) {
	mean := stats.Mean(data)
	stddev := stats.StandardDeviation(stats.Variance(data, mean))

	lower := mean - 1.5*stddev
	upper := mean + 1.5*stddev

	return int(math.Round(lower)), int(math.Round(upper))
}
