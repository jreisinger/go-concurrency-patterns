package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 5; i++ {
		// let talk whosoever is ready
		fmt.Println(<-ch)
	}
	fmt.Println("You're boring; I'm returning.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() { // just one goroutine with select
		for {
			select {
			case s := <-input1:
				ch <- s
			case s := <-input2:
				ch <- s
			}
		}
	}()
	return ch
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
