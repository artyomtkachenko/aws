package main

import "github.com/artyomtkachenko/aws/subsum"

func main() {
	data := []int{5, 5, 15, 2, 3, 10}
	sum := 15
	res := subsum.GetResult(data, sum)
	for _, el := range res {
		subsum.PrintResult(el, sum)
	}
}
