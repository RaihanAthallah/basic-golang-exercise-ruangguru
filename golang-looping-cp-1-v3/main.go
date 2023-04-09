package main

import "fmt"

func CountingNumber(n int) float64 {
	var total float64 = 0
	var i float64 = 1
	for i <= float64(n) {
		total += float64(i)
		i += 0.5
	}
	
	return total // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
