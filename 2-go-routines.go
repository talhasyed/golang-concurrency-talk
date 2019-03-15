package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateGoRoutineFibs() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 20; i++ {
		n := randomInt(39, 41)
		go Fibonacci(n)
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%s", i, n, "-"))
	}
}
