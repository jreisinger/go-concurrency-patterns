package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	str  string
	wait chan bool
}

func main() {
	ch := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-ch
		fmt.Println(msg1.str)
		msg2 := <-ch
		fmt.Println(msg2.str)

		// Each speaker must wait for a go-ahead.
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring; I'm returning.")
}

func fanIn(inputs ...<-chan message) <-chan message {
	ch := make(chan message)
	for i := range inputs {
		input := inputs[i] // new instance of 'input' for each loop
		go func() {
			for {
				ch <- <-input
			}
		}()
	}
	return ch
}

func boring(msg string) <-chan message {
	ch := make(chan message)
	waitForIt := make(chan bool) // shared between all messages
	go func() {
		for i := 0; ; i++ {
			ch <- message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return ch
}
