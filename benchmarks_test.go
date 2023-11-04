package main

import "testing"

func BenchmarkCalcFibsSequential(b *testing.B) {
	calcFibsSequential()
}

func BenchmarkCalcFibsGoRoutine(b *testing.B) {
	calcFibsGoRoutine()
}

func BenchmarkCalcFibChannels(b *testing.B) {
	calcFibChannels()
}

func BenchmarkCalcFibBlockingChannels(b *testing.B) {
	calcFibBlockingChannels()
}
