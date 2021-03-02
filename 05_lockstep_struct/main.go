// Let's put sequencing back but in different way. Each speaker must wait for a
// go-ahead. Implement this using message struct with a message string and wait
// channel.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}

type message struct {
	str  string
	wait chan bool
}

func fanIn(inputs ...<-chan message) <-chan message {
	c := make(chan message)
	for i := range inputs {
		input := inputs[i] // new instance of 'input' for each loop
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

func boring(msg string) <-chan message {
	c := make(chan message)
	waitForIt := make(chan bool) // shared between all messages
	go func() {
		for i := 0; ; i++ {
			c <- message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}
