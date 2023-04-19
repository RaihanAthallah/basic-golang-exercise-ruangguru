package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	arrLen := len(arr)
	// var newArray interface{}
	if position == "right" {
		return arr[0:(arrLen - 1)]
	}
	if position == "left" {
		return arr[1:]
	}

	return nil
}

func main() {
	dataInt := []any{1, 2, 3, 4, 5}
	dataString := []any{"Edo", "Budi", "Joko", "Tono"}

	fmt.Println(removeLeftRight(dataInt, "right"))
	fmt.Println(removeLeftRight(dataString, "left"))
}
