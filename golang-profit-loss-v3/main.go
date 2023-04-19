package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := string(content)

	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return []string{}, nil
	}
	if lines[0] == "" && len(lines) == 1 {
		return []string{}, nil
	}

	// fmt.Println(string(content))
	return lines, err // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	profit := 0
	lastDate := ""
	var splitData []string
	for _, v := range data {
		fmt.Println(v)
		splitData = strings.Split(v, ";")
		amount, _ := strconv.Atoi(splitData[2])

		lastDate = splitData[0]
		if splitData[1] == "income" {
			profit += amount
		} else if splitData[1] == "expense" {
			profit -= amount
		}
	}
	if profit >= 0 {
		return fmt.Sprintf("%s;profit;%d", lastDate, profit)
	} else {
		return fmt.Sprintf("%s;loss;%d", lastDate, -profit)
	}
	// fmt.Println(splitData[0])

	return "" // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
