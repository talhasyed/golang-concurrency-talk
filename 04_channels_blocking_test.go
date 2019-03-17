package main

import "testing"

func BenchmarkCalcFibChannelsWithSelect(b *testing.B) {
	calcFibChannelsWithSelect()
}
