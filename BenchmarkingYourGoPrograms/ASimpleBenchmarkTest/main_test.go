package main

import "testing"

func benchmarkCalculate(input int, b *testing.B) {
    for n := 0; n < b.N; n++ {
        Calculate(input)
    }
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }