package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/artyomtkachenko/aws/points" //main logic
)

func main() {
	rand.Seed(time.Now().UnixNano())
	queue := make(chan points.Points) //This channel will be used to send an endless stream of points

	k := 2                        // K closest poins
	go points.Generate(queue, 20) // Points generator. It will work 20 milliseconds and exit.

	res := points.Reduce(queue, k) // Generates the result
	fmt.Printf("Result: %+v\n", res)
}
