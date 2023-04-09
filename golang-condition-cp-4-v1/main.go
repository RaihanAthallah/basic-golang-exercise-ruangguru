package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	totalVIP := VIP * 30
	totalRegular := regular * 20
	totalStudent := student * 10
	totalPrice := totalVIP + totalRegular + totalStudent
	var fixPrice float32
	if totalPrice >= 100 {
		if day % 2 == 0 {
			if VIP+regular+student >= 5 {
				fixPrice = float32(totalPrice) * 0.80
			} else {
				fixPrice = float32(totalPrice) * 0.90
			}
		} else {
			if VIP+regular+student >= 5 {
				fixPrice = float32(totalPrice) * 0.75
			} else {
				fixPrice = float32(totalPrice) * 0.85
			}
		}
	}else{
		fixPrice = float32(totalPrice)
	}
	return fixPrice // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
