Sure! Here's a README file for your `guess-it-1` project:

---

# Guess-It-1

## Overview

`guess-it-1` is a Go-based project designed to predict the range of the next number in a series based on statistical calculations. The project uses a sliding window approach to consider the most recent numbers and computes the mean and standard deviation to predict the range in which the next number is likely to fall.

## Project Structure

The project is organized into several modules:

- `main.go`: The entry point of the application that reads input from stdin and outputs the predicted range.
- `file/file.go`: Contains functions for future file-based data input (currently not used in the main application).
- `stats/`: Implements statistical functions to compute mean, median, variance, standard deviation, and predict the range.
- `script.sh`: A shell script to run the Go program.

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://learn.zone01kisumu.ke/git/steoiro/guess-it-1.git
   cd guess-it-1
   ```

2. **Build and Run**

   Make sure you have Go installed on your system. Build and run the project using:

   ```sh
   go run main.go
   ```

## Usage

The application reads numbers from the standard input (`stdin`). For each number input, it calculates and prints the predicted range for the next number based on the last N numbers (default is 5).

### Example

```sh
$ echo -e "189\n113\n121\n114\n145\n110" | go run main.go
189 --> the standard input
120 200    --> the range for the next input, in this case for the number 113
113 --> the standard input
160 230    --> the range for the next input, in this case for the number 121
121 --> the standard input
110 140    --> the range for the next input, in this case for the number 114
114 --> the standard input
100 200    --> the range for the next input, in this case for the number 145
145 --> the standard input
1 99      --> the range for the next input, in this case for the number 110
```

## Code Structure

### `main.go`

Handles reading input and calling statistical functions to predict the range for the next number.

### `file/file.go`

Prepares for future extension to handle file-based inputs.

### `stats/`

Contains functions for statistical calculations:
- `Mean`: Calculates the average of a slice of numbers.
- `Median`: Computes the median value of the numbers.
- `Variance`: Computes the variance based on the mean.
- `StandardDeviation`: Computes the standard deviation from the variance.
- `PredictRange`: Predicts the range of the next number using mean ± 2 * standard deviation.

### `script.sh`

A shell script for running the Go program. Ensure it is executable:

```sh
chmod +x student/script.sh
```

## Testing

The program is designed to be tested with various data sets. To ensure accurate predictions, the program computes the range based on the statistical calculations and adjusts to new inputs dynamically.

## Contributing

If you would like to contribute to the project, please fork the repository and submit a pull request. We welcome improvements and additional features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For any questions or issues, please get in touch with [Stella Oiro](https://github.com/Stella-Achar-Oiro).
