package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Error("Abs(-1) = %w; want 1", got)
	}
}

func BenchmarkAbs(t *testing.B) {
	for i := 0; i < 1000; i++ {
		Abs(rand.Float64())
	}
}

func Abs(i float64) int {
	return int(math.Abs(i))
}
