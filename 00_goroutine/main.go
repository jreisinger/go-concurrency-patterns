// Launch boring goroutine (an independently executing function) and let it
// print stuff for 2 seconds. There's no communication here.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("boring")
	time.Sleep(time.Second * 2)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
