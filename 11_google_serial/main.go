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

// Google invokes Web, Image and Video searches for query serially.
func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
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
