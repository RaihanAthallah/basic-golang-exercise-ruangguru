package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	counter := 0
	for _, letter := range text {
		if letter == 'R' || letter == 'S' || letter == 'T' || letter == 'Z' || letter == 'r' || letter == 's' || letter == 't' || letter == 'z' {
			counter++
		}
	}
	return counter // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
