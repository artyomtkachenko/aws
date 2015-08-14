package points

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//ByPoints - just a slice of poins, used in sorting
type ByPoints []Points

//Points - typically a collection of points
type Points struct {
	X, Y, D int
}

func (p ByPoints) Len() int {
	return len(p)
}

func (p ByPoints) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPoints) Less(i, j int) bool {
	return p[i].D < p[j].D
}

func genRand() Points {
	n := 100
	p := new(Points)
	p.X = rand.Intn(n)
	p.Y = rand.Intn(n)

	if p.X == 0 || p.Y == 0 {
		*p = genRand()
	}
	return *p
}

//Generate generates random points and send it to the stream
func Generate(queue chan Points, howLong time.Duration) {
	quit := make(chan bool)
	timer := time.NewTimer(time.Millisecond * howLong)
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

//Reduce consumes poins from the stream and finds the K closest points.
func Reduce(queue chan Points, k int) ByPoints {
	var arr ByPoints

	for p := range queue {
		p.D = p.X*p.X + p.Y*p.Y
		fmt.Printf("Got: %+v\n", p)
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
