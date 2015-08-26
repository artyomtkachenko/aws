package main

import "fmt"

func debug(arr [][]bool) {
	for i, el := range arr {
		fmt.Println(i, el)
	}
}

func dynamic(set []int, sum int) {
	size := len(set)

	subset := make([][]bool, size)

	for i := 0; i <= size; i++ {
		subset[i] = make([]bool, sum+1) // we need it to start from 0
		subset[i][0] = true
		for j := 1; j <=sum; j++ {
			
		}
	}
	for i := 1; i <= sum; i++ {
		// subset[i] = make([]bool, size+1)
		// subset[i][0] = false
	}

	for i := 1; i <= sum; i++ {
		for j := 1; j <= size; j++ {
			subset[i][j] = subset[i][j-1]
			if i >= set[j-1] {
				subset[i][j] = subset[i][j] || subset[i-set[j-1]][j-1]
			}
		}
	}
	debug(subset)
}

func main() {
	data := []int{5, 5, 15, 10}
	sum := 15
	// _ = subsum.GetResult(data, sum)
	// for _, el := range res {
	// 	subsum.PrintResult(el, sum)
	// }
	dynamic(data, sum)
}
