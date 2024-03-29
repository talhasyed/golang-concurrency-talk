package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcFibsGoRoutine() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		n := randInt(39, 41)
		go Fib(n)
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%s", i+1, n, "-"))
	}
}
