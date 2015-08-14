package subsum

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintResult(res []int, sum int) {
	var strs []string
	for _, el := range res {
		strs = append(strs, strconv.Itoa(el))
	}
	fmt.Println(strings.Join(strs, " + "), " = ", sum)
}

func sumIt(slice []int, res []int, i int, sum int, finalResult *[][]int) int {
	if sum == 0 {

		*finalResult = append(*finalResult, res)
		return 0
	}

	if i == 0 {
		return 0
	}

	if slice[i-1] <= sum {
		sumIt(slice, res, i-1, sum, finalResult)
		res = append(res, slice[i-1])
		sumIt(slice, res, i-1, sum-slice[i-1], finalResult)
	} else {
		sumIt(slice, res, i-1, sum, finalResult)
	}
	return 0
}

func GetResult(data []int, sum int) [][]int {
	var (
		result [][]int // Final result
		res    []int   //Intermidiate result
	)
	sumIt(data, res, len(data), sum, &result)
	return result
}
