package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date < transactions[j].Date
	})

	lastDate := transactions[0].Date
	money := 0
	output := make([]string, 0)

	for _, transaction := range transactions {
		if transaction.Date == lastDate {
			if transaction.Type == "income" {
				money += transaction.Amount
			} else if transaction.Type == "expense" {
				money -= transaction.Amount
			}
		} else {
			if money > 0 {
				output = append(output, fmt.Sprintf("%s;income;%d", lastDate, money))
			} else {
				output = append(output, fmt.Sprintf("%s;expense;%d", lastDate, money*-1))
			}
			money = 0
			if transaction.Type == "income" {
				money += transaction.Amount
			} else if transaction.Type == "expense" {
				money -= transaction.Amount
			}
			lastDate = transaction.Date

		}

	}

	if money > 0 {
		output = append(output, fmt.Sprintf("%s;income;%d", lastDate, money))
	} else {
		output = append(output, fmt.Sprintf("%s;expense;%d", lastDate, money*-1))
	}

	fmt.Println(output)

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(strings.Join(output, "\n"))

	if err2 != nil {
		log.Fatal(err2)
	}

	// lastDate := transactions[0].Date
	// for i := 0; i < len(transactions); i++ {
	// 	// for i, transaction := range transactions {
	// 	income := 0
	// 	expense := 0
	// 	detail := ""

	// 	for _, t := range transactions {
	// 		if transactions[i].Date == t.Date {
	// 			if t.Type == "income" {
	// 				income += t.Amount
	// 			} else if t.Type == "expense" {
	// 				expense += t.Amount
	// 			}
	// 		}
	// 	}
	// 	total := income - expense
	// 	if total > 0 {
	// 		detail = "income"
	// 	} else if total < 0 {
	// 		total *= -1
	// 		detail = "expense"
	// 	}
	// 	if transactions[i] == transactions[i+1] {
	// 		continue
	// 	} else {
	// 		fmt.Println(transactions[i].Date, detail, total)
	// 	}
	// 	// fmt.Println(transaction.Date, detail, total)
	// }

	return nil // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"02/01/2021", "expense", 30000},
		{"02/01/2021", "income", 20000},
	}
	// var transactions2 = []Transaction{
	// 	{"01/01/2021", "income", 100000},
	// 	{"01/01/2021", "expense", 50000},
	// 	{"01/01/2021", "expense", 30000},
	// 	{"01/01/2021", "income", 20000},
	// 	{"01/01/2021", "expense", 50000},
	// }

	err := RecordTransactions("./transactions.txt", transactions)
	// err = RecordTransactions("./transactions.txt", transactions2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
