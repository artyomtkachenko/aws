package subsum

import (
	"fmt"
	"strconv"
	"strings"
)

//PrintResult prins final result
func PrintResult(res []int, sum int) {
	var strs []string
	for _, el := range res {
		strs = append(strs, strconv.Itoa(el))
	}
	fmt.Println(strings.Join(strs, " + "), " = ", sum)
}

// Recersivly search for all possible sums
func sumIt(slice []int, res []int, i int, sum int, finalResult *[][]int) int {
	if sum == 0 {

		*finalResult = append(*finalResult, res)
		return 0
	}

	if i == 0 {
		return 0
	}
	last := slice[i-1]
	if last <= sum {
		sumIt(slice, res, i-1, sum, finalResult)
		res = append(res, last)
		sumIt(slice, res, i-1, sum-last, finalResult)
	} else {
		sumIt(slice, res, i-1, sum, finalResult)
	}
	return 0
}

//GetResult invokes the main recursive function
func GetResult(data []int, sum int) [][]int {
	var (
		result [][]int // Final result
		res    []int   //Intermidiate result
	)
	sumIt(data, res, len(data), sum, &result)
	return result
}
