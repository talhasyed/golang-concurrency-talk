package main

// FibCalcChannel invokes the Fib function and adds the results to a channel
func FibCalcChannel(n int, c chan int) {
	result := Fib(n)
	c <- result
}

// FibCalcChannel2 invokes the Fib function and adds FibCalcJobResultResults to a channel
func FibCalcChannel2(id int, value int, c chan FibCalcJobResult) {
	result := Fib(value)
	c <- FibCalcJobResult{id, value, result}
}
