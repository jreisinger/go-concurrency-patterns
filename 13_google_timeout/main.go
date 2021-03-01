package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

type Result string

// Google invokes Web, Image and Video searches for query concurrently.
func Google(query string) (results []Result) {
	ch := make(chan Result)
	go func() { ch <- Web(query) }()
	go func() { ch <- Image(query) }()
	go func() { ch <- Video(query) }()

	// Let's speed things further by giving up waiting after 80 ms. But we
	// will sometimes miss some kind of search.
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// Various search kinds defined as functions.
var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}
