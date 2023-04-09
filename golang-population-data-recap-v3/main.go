package main

import (
	"fmt"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, v := range data {
		splited := strings.Split(v, ";")
		// detail := 
		for _, detail := range splited {
			fmt.Println(detail)
			
		}
		result = append(result, map[string]interface{}{})
	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi;23;Jakarta;;",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;165.42;",
	}

	fmt.Println(PopulationData(data))
}
