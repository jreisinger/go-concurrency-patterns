package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	ch := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-ch)
	}
	quit <- "Bye!"
	// Wait for Joe to let us know he's done.
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case ch <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return ch
}

func cleanup() {}
