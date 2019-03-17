package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateChannelsFibs() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int, 20)

	for i := 1; i < 20; i++ {
		n := randomInt(39, 41)
		go FibonacciChannel(i, n, c)
		result := <-c
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", i, n, result))
	}
}
