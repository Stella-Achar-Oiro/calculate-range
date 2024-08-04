package main

import (
	"bufio"
	"flag"
	"fmt"
	"guess-it-1/student/stats"
	"os"
	"strconv"
)

func main() {
    windowSize := flag.Int("window", 5, "Number of last numbers to consider for prediction")
    flag.Parse()

    var data []float64
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        value, err := strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            fmt.Println("Error")
            continue
        }
        data = append(data, value)

        if len(data) >= *windowSize {
            data = data[len(data)-*windowSize:]
        }

        if len(data) > 1 {
            lower, upper := stats.PredictRange(data)
            fmt.Printf("%d %d\n", int(lower), int(upper))
        }
    }
}

