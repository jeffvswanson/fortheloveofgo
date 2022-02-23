package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

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
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Add(%v, %v)", tc.a, tc.b), func(t *testing.T) {
			got := calculator.Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 20
	got := calculator.Multiply(5, 4)
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}
