package main

import (
	"fmt"
)

func AddElement(data []int, newData int, position string) []int {
	slice := data
	if position == "up" {
		slice = append([]int{newData}, slice...)
	}
	if position == "down" {
		slice = append(slice, newData)
	}

	return slice
}

func main() {
	slices := []int{1, 2, 3, 4, 5}
	fmt.Println(AddElement(slices, 6, "up"))
	fmt.Println(AddElement(slices, 6, "down"))
}
