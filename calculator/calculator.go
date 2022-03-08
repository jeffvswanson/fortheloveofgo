// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes numbers and returns the result of adding them all together.
func Add(a ...float64) float64 {
	var sum float64
	if len(a) == 0 {
		return sum
	}
	for _, val := range a {
		sum += val
	}
	return sum
}

// Subtract takes numbers and returns the result of subtracting the second
// from the first, the third from that result, and so on...
func Subtract(a ...float64) float64 {
	var diff float64
	if len(a) == 0 {
		return diff
	}
	diff = a[0]
	if len(a) > 1 {
		for _, val := range a[1:] {
			diff -= val
		}
	}
	return diff
}

// Multiply takes numbers and returns the result of multiplying them together.
func Multiply(a ...float64) float64 {
	var res float64
	if len(a) == 0 {
		return res
	}
	res = 1
	for _, val := range a {
		res *= val
	}
	return res
}

// Divide takes numbers and returns the result of dividing them one after another or raises an error
// when attempting to divide by 0.
func Divide(a ...float64) (float64, error) {
	var div float64
	if len(a) == 0 {
		return div, nil
	}
	div = a[0]
	if len(a) == 1 {
		return div, errors.New("at least two values are needed for division")
	}
	for _, val := range a[1:] {
		if val == 0 {
			return 0, errors.New("cannot divide by 0")
		}
		div /= val
	}
	return div, nil
}

// Sqrt takes a positive number and returns its square root or raises an error when
// attempting to take the square root of a negative number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return -1, errors.New("cannot take the square root of a negative number")
	}
	return math.Sqrt(a), nil
}
