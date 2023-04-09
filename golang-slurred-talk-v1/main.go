package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	kata := *words
	slurred := strings.Replace(kata, "s", "l", -1)
	slurred = strings.Replace(slurred, "S", "L", -1)
	slurred = strings.Replace(slurred, "r", "l", -1)
	slurred = strings.Replace(slurred, "R", "L", -1)
	slurred = strings.Replace(slurred, "z", "l", -1)
	slurred = strings.Replace(slurred, "Z", "L", -1)

	*words = slurred
	// fmt.Println(slurred)
	// TODO: answer here
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
