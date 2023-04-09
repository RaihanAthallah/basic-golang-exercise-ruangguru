package main

func SchedulableDays(date1 []int, date2 []int) []int {
	result := []int{}

	for _, v := range date1 {
		for _, v2 := range date2 {
			if v == v2 {
				result = append(result, v)
			}
		}
	}

	return result // TODO: replace this
}
