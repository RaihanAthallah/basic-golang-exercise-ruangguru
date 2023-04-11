package main

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"

	"github.com/onsi/ginkgo/v2/ginkgo/labels"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)
	for _, d := range data {
		tokens := strings.Split(d, "-")
		label := tokens[0]
		index, _ := strconv.Atoi(tokens[1])
		firstOrLast := tokens[2]
		value := tokens[3]

		
	}
	return nil // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
