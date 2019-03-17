package main

import "testing"

func BenchmarkCalcFibChannels(b *testing.B) {
	calcFibChannels()
}
