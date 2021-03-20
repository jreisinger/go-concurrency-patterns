// Get two channels from boring and print from them in lockstep (sequence).
// Channels serve as a handle on a service. We can have more instances of the
// service.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ann := boring("Ann")
	joe := boring("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
