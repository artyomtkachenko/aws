package points

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type byPoints []Points

type Points struct {
	X, Y, D int
}

func (p byPoints) Len() int {
	return len(p)
}

func (p byPoints) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byPoints) Less(i, j int) bool {
	return p[i].D < p[j].D
}

func genRand() Points {
	n := 100
	p := new(Points)
	p.X = rand.Intn(n)
	p.Y = rand.Intn(n)
	p.D = p.X*p.X + p.Y*p.Y
	if p.X == 0 || p.Y == 0 {
		*p = genRand()
	}
	return *p
}

func Generate(queue chan Points) {
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
			close(queue)
			return
		default:
			queue <- genRand()
		}
	}
}

func Reduce(queue chan Points, k int) byPoints {
	var arr byPoints

	for p := range queue {
		// fmt.Printf("Got: %+v\n", p)
		if len(arr) < k {
			arr = append(arr, p)
		} else {
			sort.Sort(arr)
			last := arr[k-1]
			if p.D <= last.D {
				arr = append(arr[:k-1], p) // Replacing the max one
			}
		}
	}
	return arr
}
