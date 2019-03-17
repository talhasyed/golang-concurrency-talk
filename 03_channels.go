package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcFibChannels() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int, 20)

	for i := 1; i < 20; i++ {
		n := randInt(39, 41)
		go FibCalcChannel(n, c)
		result := <-c
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", i, n, result))
	}
}
