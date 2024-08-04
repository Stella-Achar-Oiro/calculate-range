package file

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// ReadDataFromFile reads data from the specified file path.
func ReadDataFromFile(filePath string) ([]float64, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var data []float64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if len(line) == 0 {
            continue // Skip empty lines
        }
        num, err := strconv.ParseFloat(line, 64)
        if err != nil {
            return nil, fmt.Errorf("non-numeric value found: %s", line)
        }
        data = append(data, num)
    }
    if len(data) == 0 {
        return nil, fmt.Errorf("no valid data points: the file is empty or contains only non-numeric values")
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return data, nil
}

