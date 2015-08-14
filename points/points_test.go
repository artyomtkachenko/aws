package points

import (
	"testing"

	"github.com/artyomtkachenko/aws/points"
)

func TestReduce(t *testing.T) {
	queue := make(chan points.Points)
	k := 2
	go func() {
		queue <- points.Points{2, 2, 0}
		queue <- points.Points{2, 3, 0}
		queue <- points.Points{1, 2, 0}
		close(queue)

	}()
	res := points.Reduce(queue, k)
	t.Error(res)

}
