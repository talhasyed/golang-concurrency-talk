package main

// FibCalcChannel is a naive recursive implementation of the fibonacci algorithm
func FibCalcChannel(n int, c chan int) {
	result := Fib(n)
	c <- result
}
