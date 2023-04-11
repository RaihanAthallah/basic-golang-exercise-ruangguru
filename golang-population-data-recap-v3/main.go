package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, v := range data {
		splited := strings.Split(v, ";")
		temp := make(map[string]interface{})
		temp["name"] = splited[0]
		temp["age"], _ = strconv.Atoi(splited[1])
		temp["address"] = splited[2]
		if splited[3] != "" {
			temp["height"], _ = strconv.ParseFloat(splited[3], 64)
		}
		if splited[4] != "" {
			temp["isMarried"], _ = strconv.ParseBool(splited[4])
		}
		
		result = append(result, temp)

	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi;23;Jakarta;;",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;165.42;",
	}

	output := PopulationData(data)
	switch val := output[0]["age"].(type) {
	case int:
		fmt.Println(val + 1)
	}

}
