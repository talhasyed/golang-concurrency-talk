package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calcFibChannelsWithSelect() {
	rand.Seed(time.Now().UnixNano())
	NumCalcs := 20
	c := make(chan FibCalcJob, 20)

	for i := 1; i < NumCalcs; i++ {
		n := randInt(39, 41)
		go FibCalcChannel2(i, n, c)
	}

	for j := 1; j < NumCalcs; j++ {
		result := <-c
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", result.id, result.value, result.result))
	}
	close(c)
}
