package util

import (
	"math"
)

// Float64Compare compare float64 with accuracy
// return : 1(a>b), -1(a<b), 0(a=b)
func Float64Compare(a, b float64, accuracy func() float64) int {
	switch delta := math.Abs(a - b); {
	case delta < accuracy() || a-b == 0:
		return 0
	case a-b > 0:
		return 1
	default:
		return -1
	}
}

// Float64AccuracyDefault default accuracy 0.00001
func Float64AccuracyDefault() float64 {
	return 0.00001
}

// CompareFloat64 compare a & b with default accuracy(Float64AccuracyDefault())
func CompareFloat64(a, b float64) int {
	return Float64Compare(a, b, Float64AccuracyDefault)
}
