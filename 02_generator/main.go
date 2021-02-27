package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := boring("boring!") // get channel from a function
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("You're boring; I'm returning.")
}

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() { // launch the goroutine from inside a function
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch // return the channel
}
