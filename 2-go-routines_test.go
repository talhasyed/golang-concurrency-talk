package main

import "testing"

func BenchmarkCalculateGoRoutineFibs(b *testing.B) {
	calculateGoRoutineFibs()
}
