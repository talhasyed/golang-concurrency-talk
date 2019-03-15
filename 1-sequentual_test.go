package main

import "testing"

func BenchmarkFib10SequentialFibs(b *testing.B) {
	calculateSequentialFibs()
}
