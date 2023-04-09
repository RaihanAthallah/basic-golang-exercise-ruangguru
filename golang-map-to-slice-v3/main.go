package main

func MapToSlice(mapData map[string]string) [][]string {
	slices := make([][]string, 0)
	for key, value := range mapData {
		slices = append(slices, []string{key, value})
	}
	// fmt.Println(slices)
	return slices // TODO: replace this
}

func main() {
	mapData := map[string]string{
		"hello": "world",
		"John":  "Doe",
		"age":   "14",
	}
	MapToSlice(mapData)
	// result := MapToSlice()
}
