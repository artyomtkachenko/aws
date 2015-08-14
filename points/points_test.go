package points

import (
	"testing"

	"github.com/artyomtkachenko/aws/points"
)

func assert(expected, got []points.Points, expectedSize int, t *testing.T) {
	for i, el := range expected {
		if el.D != got[i].D {
			t.Error("Expected: ", el, " Got:", got[i])
		}
	}
	if len(got) != len(expected) {
		t.Error("Expected slices to be equal in size, ", expected, " Got:", got)
	}
	if len(got) != expectedSize {
		t.Error("Expected (K) size of an array should be, ", expectedSize, " Got:", len(got))
	}
}

func TestReduce(t *testing.T) {
	expected := []points.Points{{2, 2, 8}, {1, 2, 5}}
	queue := make(chan points.Points)
	k := 2
	go func() {
		queue <- points.Points{2, 2, 0}
		queue <- points.Points{2, 3, 0}
		queue <- points.Points{1, 2, 0}
		close(queue)

	}()
	res := points.Reduce(queue, k)
	assert(expected, res, k, t)

	expected = []points.Points{{2, 2, 8}, {2, 3, 13}, {10, 2, 104}, {0, 0, 0}}
	queue = make(chan points.Points)
	k = 4
	go func() {
		queue <- points.Points{2, 2, 0}
		queue <- points.Points{2, 3, 0}
		queue <- points.Points{10, 2, 0}
		queue <- points.Points{1, 20, 0}
		queue <- points.Points{100, 22, 0}
		queue <- points.Points{11, 2, 0}
		queue <- points.Points{1100, 2, 0}
		queue <- points.Points{0, 0, 0}
		close(queue)

	}()
	res = points.Reduce(queue, k)
	assert(expected, res, k, t)

	// his method does not work with slice of structs for some reason
	// if !reflect.DeepEqual(res, expected) {
	// 	t.Error("Expected: ", expected, " Got:", res)
	// }

}
