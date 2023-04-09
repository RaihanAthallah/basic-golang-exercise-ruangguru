package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	serial := strconv.Itoa(numbers)
	biggest := 0
	pair := 0
	// fmt.Printf(serial)
	for i := 1; i < len(serial); i++ {

		num1, _ := strconv.Atoi(string(serial[i]))
		num2, _ := strconv.Atoi(string(serial[i-1]))
		if num1+num2 > biggest {
			biggest = num1 + num2

			pairNumb, _ := strconv.Atoi(string(serial[i-1]) + string(serial[i]))
			pair = pairNumb
		}
	}

	return pair // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
