package main

import (
	"fmt"
)

func GetPredicate(math, science, english, indonesia int) string {
	avg := (math + science + english + indonesia) / 4
	message := ""
	if avg == 100 {
		message = "Sempurna"
	}else if avg >= 90 {
		message = "Sangat Baik"
	}else if avg >= 80 {
		message = "Baik"
	}else if avg >= 70 {
		message = "Cukup"
	}else if avg >= 60 {
		message = "Kurang"
	}else if avg < 60 {
		message = "Sangat kurang"
	}else{
	 // TODO: replace this
	}
	return message
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
