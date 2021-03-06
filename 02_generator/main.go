// This time get channel from boring instead of creating it in main. And launch
// the goroutine from inside boring instead of main.
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
