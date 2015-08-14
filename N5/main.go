package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/artyomtkachenko/aws/points"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	queue := make(chan points.Points) // Buffered channel for a stream

	k := 2
	go points.Generate(queue)

	res := points.Reduce(queue, k)
	fmt.Printf("Result: %+v\n", res)
}
