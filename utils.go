package main

import (
	"math/rand"
)

// randomInt returns a random intger between min and max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Fibonacci is a naive recursive implementation of the fibonacci algorithm
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciChannel is a naive recursive implementation of the fibonacci algorithm
func FibonacciChannel(worker int, n int, c chan int) {
	result := Fibonacci(n)
	c <- result
}
