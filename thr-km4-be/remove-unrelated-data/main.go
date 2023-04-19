package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	m := make(map[string]any)
	m["name"] = "Edo"
	m["age"] = 20
	m["address"] = "Jakarta"

	fmt.Println(removeUnrelated(m, "address"))
}
