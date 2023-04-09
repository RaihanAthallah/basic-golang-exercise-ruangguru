package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {

	word := ""
	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		}
		if str[i] == ' ' || i == len(str)-1 {
			reverved := ""
			for j := len(word) - 1; j >= 0; j-- {
				reverved += string(word[j])
			}
			if unicode.IsUpper(rune(word[0])) == true {
				reverved = strings.ToUpper(string(reverved[0])) + reverved[1:]
			}
			if unicode.IsLower(rune(word[len(word)-1])) == true {
				reverved = reverved[:len(reverved)-1] + strings.ToLower(string(reverved[len(reverved)-1]))
			}
			result += reverved + " "
			word = ""
		}

	}

	return result[:len(result)-1] // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
