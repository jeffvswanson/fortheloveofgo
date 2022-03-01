package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"testing"
)

// Tests floating point numbers to ensure parity due to rounding errors
func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{1, 1, 2},
		{1, 0, 1},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, 0, -1},
		{-1, -1, -2},
		{0.1, 0.9, 1},
		{-0.1, -0.9, -1},
		{-0.1, -0.1, -0.2},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Add(%v, %v)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Add(tc.a, tc.b)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, -1},
		{0, 0, 0},
		{-1, -1, 0},
		{0.1, 0.1, 0},
		{-1.1, -0.2, -0.9},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Subtract(%v. %v)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Subtract(tc.a, tc.b)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{1, 1, 1},
		{0, 1, 0},
		{1, 0, 0},
		{1, 0.1, 0.1},
		{-1, 1, -1},
		{-1, -1, 1},
		{-1, 0.1, -0.1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Multiply(%v, %v)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Multiply(tc.a, tc.b)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("want %f, got %f", tc.want, got)
			}

		})
	}
}
