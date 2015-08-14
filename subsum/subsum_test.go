package subsum

import (
	"reflect"
	"testing"

	"github.com/artyomtkachenko/aws/subsum"
)

func TestGetResult(t *testing.T) {

	var expected [][]int
	data := []int{5, 5, 15, 2, 3, 10}
	sum := 15
	res := subsum.GetResult(data, sum)
	expected = append(expected, []int{15})
	expected = append(expected, []int{3, 2, 5, 5})
	expected = append(expected, []int{10, 5})
	expected = append(expected, []int{10, 5})
	expected = append(expected, []int{10, 3, 2})

	if !reflect.DeepEqual(res, expected) {
		t.Error("Expected: ", expected, " Got:", res)
	}
	// Negative testing
	expected = append(expected, []int{100, 3, 2})
	if reflect.DeepEqual(res, expected) {
		t.Error("Expected: ", expected, " Got:", res)
	}
}
