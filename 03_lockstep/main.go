package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Channels as a handle on a service
	joe := boring("Joe")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		// joe and ann count in lockstep (sequence)
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring; I'm returning.")
}

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
