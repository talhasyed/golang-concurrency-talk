package main

// FibCalcChannel is a naive recursive implementation of the fibonacci algorithm
func FibCalcChannel(n int, c chan int) {
	result := Fib(n)
	c <- result
}

// FibCalcChannel2 is a naive recursive implementation of the fibonacci algorithm
func FibCalcChannel2(id int, value int, c chan FibCalcJob) {
	result := Fib(value)
	c <- FibCalcJob{id, value, result}
}
