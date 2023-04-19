package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/service"
	"fmt"
)

func CashierApp(db *database.Database) service.ServiceInterface {
	service := service.NewService(db)

	return service
}

// gunakan untuk debugging
func main() {
	database := database.NewDatabase()
	service := CashierApp(database)

	// fmt.Println(" ",service.ShowCart())
	err := service.AddCart("Kaos Polos", 2)
	// fmt.Println(service.ShowCart())
	if err != nil {
		panic(err)
	}
	// err = service.RemoveCart("Kaos Polos")
	// service.AddCart("Kaos sablon", 1)

	paymentInformation, err := service.Pay(100000)
	fmt.Println(paymentInformation)
	items, _ := database.GetCartItems()
	for _, item := range items {
		fmt.Printf("hasil %v", item)
	}
	if err != nil {
		panic(err)
	}

	// err := service.AddCart("Kaos Polos abizzsss", 2)

	if err != nil {
		panic(err)
	}

	fmt.Println("success")
}
