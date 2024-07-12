package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadDataFromFile reads a file from the given file path, and parses its content into a slice of float64 values.
func ReadDataFromFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data []float64
	scanner := bufio.NewScanner(file)
	// Scan the file line by line.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue // Skip empty lines
		}
		// Parse the line into a float64 value.
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("non-numeric value found: %s", line)
		}
		// Append the parsed value to the data slice.
		data = append(data, num)
	}
	// If the data slice is still empty after parsing, return an error.
	if len(data) == 0 {
		return nil, fmt.Errorf("no valid data points: the file is empty or contains only non-numeric values")
	}
	// Check for any scanning errors.
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file %s: %v", filePath, err)
		return nil, err
	}
	// Return the parsed data slice and no error.
	return data, nil
}
