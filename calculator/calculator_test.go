package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
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
		t.Run(fmt.Sprintf("Add(%f, %f)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Add(tc.a, tc.b)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestAddRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		t.Run(fmt.Sprintf("AddRand(%f, %f)", a, b), func(t *testing.T) {
			got := calculator.Add(a, b)
			if !withinTolerance(got, want, 1e-12) {
				t.Errorf("got %f, want %f", got, want)
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
		t.Run(fmt.Sprintf("Subtract(%f, %f)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Subtract(tc.a, tc.b)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %f, want %f", got, tc.want)
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
		t.Run(fmt.Sprintf("Multiply(%f, %f)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Multiply(tc.a, tc.b)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %f, want %f", got, tc.want)
			}

		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a           float64
		b           float64
		want        float64
		errExpected bool
	}{
		{1, 1, 1, false},
		{0, 1, 0, false},
		{0.1, 1, 0.1, false},
		{0, 0.1, 0, false},
		{-1, 1, -1, false},
		{-1, -1, 1, false},
		{-0.1, 1, -0.1, false},
		{1, 0, 0, true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Divide(%f, %f)", tc.a, tc.b), func(t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b)
			if (tc.errExpected && err == nil) || (!tc.errExpected && err != nil) {
				t.Errorf("Improper error condition, err.Expected=%t, got: '%s'", tc.errExpected, err)
			}
			if err == nil {
				if !withinTolerance(got, tc.want, 1e-12) {
					t.Errorf("got %f, want %f", got, tc.want)
				}
			}
		})
	}
}
