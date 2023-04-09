package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	// fmt.Println(products)
	var change []int = []int{}
	total := 0
	for _, product := range products {
		total += product.Price + product.Tax
	}
	fmt.Println(total)
	sisa := amount - total
	for sisa > 0 {
		// fmt.Println(sisa)
		if sisa >= 1000 {
			sisa -= 1000
			change = append(change, 1000)
		} else if sisa >= 500 {
			sisa -= 500
			change = append(change, 500)
		} else if sisa >= 200 {
			sisa -= 200
			change = append(change, 200)
		} else if sisa >= 100 {
			sisa -= 100
			change = append(change, 100)
		} else if sisa >= 50 {
			sisa -= 50
			change = append(change, 50)
		} else if sisa >= 20 {
			sisa -= 20
			change = append(change, 20)
		} else if sisa >= 10 {
			sisa -= 10
			change = append(change, 10)
		} else if sisa >= 5 {
			sisa -= 5
			change = append(change, 5)
		} else {
			sisa -= 1
			change = append(change, 1)
		}
	}
	if sisa == 0 {
		change = append(change)
	}
	return change // TODO: replace this
}

func main() {
	fmt.Println(MoneyChanges(10000, []Product{
		{Name: "Baju", Price: 5000, Tax: 500},
		{Name: "Celana", Price: 3000, Tax: 300},
	}))
}

/* [{Name: "Baju",
Price: 5000,
Tax: 500},

{Name: "Celana",
Price: 3000,
Tax: 300}]

*/
