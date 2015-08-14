package subsum

import "testing"

func TestGetResult(t *testing.T) {

	var expected [][]int
	data := []int{5, 5, 15, 2, 3, 10}
	sum := 15
	_ = subsum.GetResult(data, sum)
	// expected := [][]int{[15], [3, 2, 5, 5], [10, 5] [10, 5], [10, 3, 2]}
	expected = append(expected, []int{15})
	expected = append(expected, []int{3, 2, 5, 5})
	expected = append(expected, []int{10, 5})
	expected = append(expected, []int{10, 5})
	expected = append(expected, []int{10, 3, 2})

	if sum != 16 {
		t.Error(expected)
	}
}
