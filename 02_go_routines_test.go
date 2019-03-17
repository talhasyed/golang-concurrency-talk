package main

import "testing"

func BenchmarkCalcFibsGoRoutine(b *testing.B) {
	calcFibsGoRoutine()
}
