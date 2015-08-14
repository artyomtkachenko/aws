package main

import (
	"fmt"
	"math/rand"
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
func genRand() points {
	n := 100
	p := new(points)
	p.x = rand.Intn(n)
	p.y = rand.Intn(n)
	p.d = p.x*p.x + p.y*p.y
	if p.x == 0 || p.y == 0 {
		*p = genRand()
	}
	return *p
}

func generate(queue chan points) {
	quit := make(chan bool)
	timer := time.NewTimer(time.Millisecond * 20)
	go func() {
		<-timer.C
		fmt.Println("Timer expired")
		quit <- true
	}()
	for {
		select {
		case <-quit:
			// time.Sleep(9000 * time.Millisecond) // Gi
			fmt.Println("Got QUIT")
			close(queue)
			return
		default:
			queue <- genRand()
		}
	}

}

func reduce(queue chan points, k int) {
	var arr byPoints

	for p := range queue {

		// fmt.Printf("Got: %+v\n", p)
		if len(arr) < k {
			arr = append(arr, p)
		} else {
			// fmt.Printf("Before sort Got: %+v\n", arr)
			// sort.Sort(arr)
			// fmt.Printf("After sirt Got: %+v\n", arr)
			last := arr[k-1]
			if p.d <= last.d {
				arr = append(arr[:k-1], p) // Replacing the max one
			}
		}
	}
	fmt.Printf("Got: %+v\n", arr)
	// fmt.Println(arr)
}

func main() {
	queue := make(chan points) // Buffered channel for a stream

	k := 2
	go generate(queue)
	reduce(queue, k)
}
