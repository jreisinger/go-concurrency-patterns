// Generator is a function that returns a channel. We launch the goroutine from
// inside the function.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
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
