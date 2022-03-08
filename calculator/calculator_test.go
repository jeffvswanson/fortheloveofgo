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
		a    []float64
		want float64
	}{
		{nil, 0},
		{[]float64{0}, 0},
		{[]float64{1, 1}, 2},
		{[]float64{1, 0}, 1},
		{[]float64{0, 0}, 0},
		{[]float64{-1, 1}, 0},
		{[]float64{-1, 0}, -1},
		{[]float64{-1, -1}, -2},
		{[]float64{0.1, 0.9}, 1},
		{[]float64{-0.1, -0.9}, -1},
		{[]float64{-0.1, -0.1}, -0.2},
		{[]float64{1, 2, 3}, 6},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Add(%v)", tc.a), func(t *testing.T) {
			got := calculator.Add(tc.a...)
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
		a    []float64
		want float64
	}{
		{nil, 0},
		{[]float64{0}, 0},
		{[]float64{1, 1}, 0},
		{[]float64{1, 0}, 1},
		{[]float64{0, 1}, -1},
		{[]float64{0, 0}, 0},
		{[]float64{-1, -1}, 0},
		{[]float64{0.1, 0.1}, 0},
		{[]float64{-1.1, -0.2}, -0.9},
		{[]float64{-1, -2, -3}, 4},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Subtract(%v)", tc.a), func(t *testing.T) {
			got := calculator.Subtract(tc.a...)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a    []float64
		want float64
	}{
		{nil, 0},
		{[]float64{0}, 0},
		{[]float64{1}, 1},
		{[]float64{2}, 2},
		{[]float64{1, 1}, 1},
		{[]float64{0, 1}, 0},
		{[]float64{1, 0}, 0},
		{[]float64{1, 0.1}, 0.1},
		{[]float64{-1, 1}, -1},
		{[]float64{-1, -1}, 1},
		{[]float64{-1, 0.1}, -0.1},
		{[]float64{10, 10, 10}, 1000},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Multiply(%v)", tc.a), func(t *testing.T) {
			got := calculator.Multiply(tc.a...)
			if !withinTolerance(got, tc.want, 1e-12) {
				t.Errorf("got %f, want %f", got, tc.want)
			}

		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a           []float64
		want        float64
		errExpected bool
	}{
		{nil, 0, false},
		{[]float64{1}, 1, true}, // Need at least two values for division
		{[]float64{1, 1}, 1, false},
		{[]float64{0, 1}, 0, false},
		{[]float64{0.1, 1}, 0.1, false},
		{[]float64{0, 0.1}, 0, false},
		{[]float64{-1, 1}, -1, false},
		{[]float64{-1, -1}, 1, false},
		{[]float64{-0.1, 1}, -0.1, false},
		{[]float64{1000, 10, 10}, 10, false},
		{[]float64{1, 0}, 0, true},    // Division by 0
		{[]float64{4, 2, 0}, 0, true}, // Division by 0
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Divide(%v)", tc.a), func(t *testing.T) {
			got, err := calculator.Divide(tc.a...)
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

func TestSqrt(t *testing.T) {
	tests := []struct {
		a           float64
		want        float64
		errExpected bool
	}{
		{0, 0, false},
		{1, 1, false},
		{4, 2, false},
		{2, 1.414, false},
		{-1, -1, true}, // Imaginary number
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Sqrt(%f)", tc.a), func(t *testing.T) {
			got, err := calculator.Sqrt(tc.a)
			if (tc.errExpected && err == nil) || (!tc.errExpected && err != nil) {
				t.Errorf("Improper error condidion, errExpected=%t, got: '%s'", tc.errExpected, err)
			}
			if err == nil {
				if !withinTolerance(got, tc.want, 1e-3) {
					t.Errorf("got %f, want %f", got, tc.want)
				}
			}
		})
	}
}
