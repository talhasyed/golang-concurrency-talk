package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcFibBlockingChannels() {
	rand.Seed(time.Now().UnixNano())
	NumCalcs := 20

	c := make(chan FibCalcJob, 20)

	for i := 0; i < NumCalcs; i++ {
		n := randInt(39, 41)
		go FibCalcChannel2(i+1, n, c)
	}

	for j := 0; j < NumCalcs; j++ {
		result := <-c
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", result.id, result.value, result.result))
	}
	close(c)
}
