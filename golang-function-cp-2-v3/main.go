package main

import (
	"fmt"
	"strings"
)

// var Vowel = []string{"a", "i", "u", "e","o","A","I","U","E","O"}

func CountVowelConsonant(str string) (int, int, bool) {
	str = strings.ToLower(str)
	countVowel := 0
	countNonAlphabet := 0
	check := true
	for _, s := range str {
		if strings.ContainsAny(string(s), "aiueo") {
			countVowel++
		}else if strings.ContainsAny(string(s), ", ") {
			countNonAlphabet++
		}
	}
	if countVowel > 0 &&  len(str)> 0 {
			check = false
		}
	return countVowel, len(str)-countVowel-countNonAlphabet, check // TODO: replace this
}
// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku"))
}
