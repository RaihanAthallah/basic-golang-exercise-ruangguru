package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	var splitted []string
	if strings.Contains(names, ";") {
		splitted = strings.Split(names, ";")
	} else if strings.Contains(names, " ") {
		splitted = strings.Split(names, " ")
	} else if strings.Contains(names, ",") {
		splitted = strings.Split(names, ",")
	}
	shortest := "MuhammadRaihanAthallah"
	for _, name := range splitted {
		if len(name) < len(shortest) {
			shortest = name
		} else if len(name) == len(shortest) {
			if name < shortest {
				shortest = name
			}
		}
	}
	return shortest // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
	// fmt.Println(utf8.RuneCountInString("Ari"))                            // "Tia"
	// fmt.Println(strconv.Atoi("220"))                                      // "Tia"
}
