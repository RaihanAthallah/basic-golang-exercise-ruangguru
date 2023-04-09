package main

func ExchangeCoin(amount int) []int {
	exchanged := []int{}
	for amount > 0 {
		if amount >= 1000 {
			amount -= 1000
			exchanged = append(exchanged, 1000)
		} else if amount >= 500 {
			amount -= 500
			exchanged = append(exchanged, 500)
		} else if amount >= 200 {
			amount -= 200
			exchanged = append(exchanged, 200)
		} else if amount >= 100 {
			amount -= 100
			exchanged = append(exchanged, 100)
		} else if amount >= 50 {
			amount -= 50
			exchanged = append(exchanged, 50)
		} else if amount >= 20 {
			amount -= 20
			exchanged = append(exchanged, 20)
		} else if amount >= 10 {
			amount -= 10
			exchanged = append(exchanged, 10)
		} else if amount >= 5 {
			amount -= 5
			exchanged = append(exchanged, 5)
		} else if amount >= 1 {
			amount -= 1
			exchanged = append(exchanged, 1)
		}
	}

	return exchanged // TODO: replace this
}
