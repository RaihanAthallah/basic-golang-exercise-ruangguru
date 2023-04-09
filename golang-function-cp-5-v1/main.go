package main

import (
	"fmt"
	"sort"
)

func FindMin(nums ...int) int {
	sort.Ints(nums)
	min := nums[0]
	// fmt.Println(min)
	return min // TODO: replace this
}

func FindMax(nums ...int) int {
	sort.Ints(nums)
	max := nums[len(nums)-1]
	// fmt.Println(max)
	return max // TODO: replace this
}

func SumMinMax(nums ...int) int {
	sort.Ints(nums)
	return FindMin(nums[0]) + FindMax(nums[len(nums)-1]) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
