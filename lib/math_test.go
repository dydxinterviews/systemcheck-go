package lib_test

import (
	"errors"
	"testing"

	"github.com/dydxprotocol/go-index-price/lib"
	"github.com/stretchr/testify/require"
)

const (
	zero = float64(0)
)

func TestMedian(t *testing.T) {
	median, err := lib.Median([]string{
		"1",
		"2",
	})
	require.Equal(t, median, float64(1.5))
	require.NoError(t, err)
}

func TestMedianTableTest(t *testing.T) {
	tests := map[string]struct { // This is a table test.
		// variables
		values []string

		// expectations
		expectedError  error
		expectedResult float64
	}{
		"Nonnumeric value": {
			values: []string{
				"foo",
				"1",
				"2",
			},
			expectedError: errors.New("strconv.ParseFloat: parsing \"foo\": invalid syntax"),
		},
		"0 values": {
			values:        []string{},
			expectedError: lib.NoMedianError,
		},
		"odd values": {
			values: []string{
				"1",
				"2",
				"3",
			},
			expectedResult: float64(2),
		},
		"even values": {
			values: []string{
				"1",
				"2",
				"3",
				"4",
			},
			expectedResult: float64(2.5),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			median, err := lib.Median(tc.values)
			if tc.expectedError != nil {
				require.Equal(t, zero, median)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResult, median)
			}
		})
	}
}
