package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcFibsSequential() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 20; i++ {
		n := randInt(39, 41)
		result := Fib(n)

		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", i, n, result))
	}
}
