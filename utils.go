package main

import (
	"math/rand"
)

// randInt returns a random intger between min and max
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Fib is a naive recursive implementation of the fibonacci algorithm
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
