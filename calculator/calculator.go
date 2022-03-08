// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing them or raises an error
// when attempting to divide by 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

// Sqrt takes a positive number and returns its square root or raises an error when
// attempting to take the square root of a negative number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return -1, errors.New("cannot take the square root of a negative number")
	}
	return math.Sqrt(a), nil
}
