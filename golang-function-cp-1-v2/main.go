package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	hari := ""
	bulan := ""
	if day < 10 {
		hari = strconv.Itoa(day)
		hari = "0" + hari
	} else {
		hari = strconv.Itoa(day)
	}
	if month == 1 {
		bulan = "January"
	} else if month == 2 {
		bulan = "February"
	} else if month == 3 {
		bulan = "March"
	} else if month == 4 {
		bulan = "April"
	} else if month == 5 {
		bulan = "May"
	} else if month == 6 {
		bulan = "June"
	} else if month == 7 {
		bulan = "July"
	} else if month == 8 {
		bulan = "August"
	} else if month == 9 {
		bulan = "September"
	} else if month == 10 {
		bulan = "October"
	} else if month == 11 {
		bulan = "November"
	} else if month == 12 {
		bulan = "December"
	}

	tanggal := hari + "-" + bulan + "-" + strconv.Itoa(year)
	return tanggal // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
