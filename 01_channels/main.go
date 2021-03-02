// Launch boring() via a gourotine and read its output from a channel 5 times.
// Boring() echoes back the argument in random (max 1 second) intervals. It
// appends a sequence number and sends the string via a channel.

// Do not communicate by sharing memory; instead, share memory by communicating.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I'm returning.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
