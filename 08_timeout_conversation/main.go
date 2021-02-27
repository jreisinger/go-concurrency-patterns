package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := boring("Joe")
	timeout := time.After(3 * time.Second)
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
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
