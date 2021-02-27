package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	ch := fanIn(joe, ann)

	for i := 0; i < 5; i++ {
		// let talk whosoever is ready
		fmt.Println(<-ch)
	}
	fmt.Println("You're boring; I'm returning.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-input1
		}
	}()
	go func() {
		for {
			ch <- <-input2
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
