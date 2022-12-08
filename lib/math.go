package lib

import (
	"errors"
	"sort"
	"strconv"
)

var (
	NoMedianError = errors.New("Cannot get median of zero values")
)

// Median gets the median value as a float from an array of numeric strings.
func Median(values []string) (float64, error) {
	if len(values) == 0 {
		return 0, NoMedianError
	}

	numericValues := make([]float64, len(values))
	for i, val := range values {
		numericValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, err
		}
		numericValues[i] = numericValue
	}

	sort.Float64s(numericValues) // sort the numbers

	mNumber := len(numericValues) / 2

	if len(numericValues)%2 != 0 {
		return numericValues[mNumber], nil
	}

	return (numericValues[mNumber-1] + numericValues[mNumber]) / 2, nil
}
