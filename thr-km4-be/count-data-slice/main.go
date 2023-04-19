package main

import "fmt"

func howManyElements(data []any) int {
	sliceLen := len(data)
	return sliceLen
}

func main() {
	integer := []any{1, 2, 3, 4, 5}
	stringer := []any{"Edo", "Budi", "Joko", "Tono"}
	boolean := []any{true, false, true, false, true, false}
	fmt.Println(howManyElements(integer))
	fmt.Println(howManyElements(stringer))
	fmt.Println(howManyElements(boolean))
}
