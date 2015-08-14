package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	queue := make(chan points) // Buffered channel for a stream

	k := 2
	go points.Generate(queue)
	points.Reduce(queue, k)
}
