// Do not communicate by sharing memory; instead, share memory by communicating.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// var ch chan string
	// ch = make(chan int)
	ch := make(chan string)  // declare and initialize channel in main
	go boring("boring!", ch) // launch (another) goroutine from main
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("You're boring; I'm returning.")
}

func boring(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
