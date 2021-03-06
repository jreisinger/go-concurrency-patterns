// Launch boring in a goroutine and let it talk a bit. There's no communication.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("boring!")
	time.Sleep(time.Second * 2)
	fmt.Println("I'm returning.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
