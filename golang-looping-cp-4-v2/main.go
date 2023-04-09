package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	// splitted := 
	domain := strings.Split(email, "@")[1]
	domain = strings.Split(domain, ".")[0]
	tld := ""
	for i, v := range email {
		if string(v) == "."{
			tld = email[i+1:]
			break
		}
	}
	response := fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
	return response // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
