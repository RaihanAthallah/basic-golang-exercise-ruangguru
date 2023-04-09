package main

import (
	"fmt"
)

func ProviderChecker(provider string, result *string) {
}
func PhoneNumberChecker(number string, result *string) {
	// fmt.Println(number[:2])
	// fmt.Println(number[2:3])
	if number[:2] == "62" && number[2:3] == "8" {
		if len(number) >= 11 {
			if number[3:5] == "11" || number[3:5] == "12" || number[3:5] == "13" || number[3:5] == "14" || number[3:5] == "15" {
				*result = "Telkomsel"
			} else if number[3:5] == "16" || number[3:5] == "17" || number[3:5] == "18" || number[3:5] == "19" {
				*result = "Indosat"
			} else if number[3:5] == "21" || number[3:5] == "22" || number[3:5] == "23" {
				*result = "XL"
			} else if number[3:5] == "27" || number[3:5] == "28" || number[3:5] == "29" {
				*result = "Tri"
			} else if number[3:5] == "52" || number[3:5] == "53" {
				*result = "AS"
			} else if number[3:5] == "81" || number[3:5] == "82" || number[3:5] == "83" || number[3:5] == "84" || number[3:5] == "85" || number[3:5] == "86" || number[3:5] == "87" || number[3:5] == "88" {
				*result = "Smartfren"
			} else {
				*result = "invalid"
			}
		} else {
			*result = "invalid"
		}
	} else if number[0:2] == "08" {
		if len(number) >= 10 {
			if number[2:4] == "11" || number[2:4] == "12" || number[2:4] == "13" || number[2:4] == "14" || number[2:4] == "15" {
				*result = "Telkomsel"
			} else if number[2:4] == "16" || number[2:4] == "17" || number[2:4] == "18" || number[2:4] == "19" {
				*result = "Indosat"
			} else if number[2:4] == "21" || number[2:4] == "22" || number[2:4] == "23" {
				*result = "XL"
			} else if number[2:4] == "27" || number[2:4] == "28" || number[2:4] == "29" {
				*result = "Tri"
			} else if number[2:4] == "52" || number[2:4] == "53" {
				*result = "AS"
			} else if number[2:4] == "81" || number[2:4] == "82" || number[2:4] == "83" || number[2:4] == "84" || number[2:4] == "85" || number[2:4] == "86" || number[2:4] == "87" || number[2:4] == "88" {
				*result = "Smartfren"
			} else {
				*result = "invalid"
			}
		} else {
			*result = "invalid"
		}
	} else {
		*result = "invalid"
	}

	// TODO: answer here
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081211111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
