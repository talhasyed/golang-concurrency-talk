package main

import "testing"

func BenchmarkCalcFibsSequential(b *testing.B) {
	calcFibsSequential()
}
