package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type byPoints []points

type points struct {
	x, y, d int
}

func (p byPoints) Len() int {
	return len(p)
}

func (p byPoints) Swap(i, j int) {
	p[i].d, p[j].d = p[j].d, p[i].d
}

func (p byPoints) Less(i, j int) bool {
	return p[i].d < p[j].d
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generate(queue chan points) {
	for i := 0; i < 10; i++ {
		n := 100
		p := new(points)
		p.x = rand.Intn(n)
		p.y = rand.Intn(n)
		p.d = p.x*p.x + p.y*p.y
		queue <- *p
	}
}

func reduce(queue chan points, k int) {
	var arr byPoints
	for i := 0; i < 10; i++ {
		p := <-queue
		fmt.Println("Got: ", p)
		if len(arr) < k {
			arr = append(arr, p)
		} else {
			sort.Sort(arr)
			last := arr[k-1]
			if p.d <= last.d {
				arr = append(arr[:k-1], p) // Replacing the max one
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	queue := make(chan points) // Buffered channel for a stream
	k := 2
	go generate(queue)
	reduce(queue, k)
}
