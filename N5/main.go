package main

import (
	"math/rand"
	"time"

	"github.com/artyomtkachenko/aws/points"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	queue := make(chan points.Points) // Buffered channel for a stream

	k := 2
	go points.Generate(queue)
	points.Reduce(queue, k)
}
