package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Tell Joe to stop when we're tired of listening to him.
	quit := make(chan bool)
	ch := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-ch)
	}
	quit <- true
}

func boring(msg string, quit <-chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case ch <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
		}
	}()
	return ch
}
