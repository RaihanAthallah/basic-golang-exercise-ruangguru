package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	for _, v := range data {
		if strings.Contains(v, input) {
			result += v + ","
		}

	}

	return result[:len(result)-1] // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
