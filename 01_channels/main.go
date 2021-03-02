// Launch boring() via a gourotine and read its output from a channel 5 times.
// Boring() echoes back the argument in random (max 1 second) intervals.

// Do not communicate by sharing memory; instead, share memory by communicating.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	go boring("boring!", ch)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("I'm returning.")
}

func boring(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
