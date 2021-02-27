package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := boring("boring!")
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		// time.After returns channel that blocks for a duration then
		// returns current time, once.
		case <-time.After(800 * time.Millisecond):
			fmt.Println("You are too slow.")
			return
		}
	}
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
