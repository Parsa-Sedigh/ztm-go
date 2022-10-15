package main

import "fmt"

// because of variadics, nums will be a slice of int, in this case
func sum(nums ...int) int {
	sum := 0

	for _, n := range nums {
		sum += n // accumulate number with sum variable
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	// expand all into it's consituents
	answer := sum(all...)

	fmt.Println(answer)
}
