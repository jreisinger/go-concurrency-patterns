// Launch boring goroutine and read its output from a channel 5 times.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
