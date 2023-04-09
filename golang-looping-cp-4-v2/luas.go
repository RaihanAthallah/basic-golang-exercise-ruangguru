package luas

import (
	"fmt"
)

func Luas() {
	for {
		var bentuk string
		var sisi, alas, tinggi float64

		fmt.Println("=== Kalkulator Geometri ===")

		fmt.Print("Masukkan bentuk geometri (persegi/persegi-panjang): ")
		fmt.Scanln(&bentuk)

		switch SelectForm(bentuk) {
		case "persegi":
			fmt.Print("Masukkan sisi: ")
			fmt.Scanln(&sisi)

			resultLuas, resultKeliling, err := CalculateSquare(sisi)
			if err != "" {
				fmt.Println(err)
			}

			fmt.Printf("Luas persegi: %.2f\n", resultLuas)
			fmt.Printf("Keliling persegi: %.2f\n", resultKeliling)
		case "persegi-panjang":
			fmt.Print("Masukkan panjang: ")
			fmt.Scanln(&alas)

			fmt.Print("Masukkan lebar: ")
			fmt.Scanln(&tinggi)

			resultLuas, resultKeliling, err := CalculateRectangle(alas, tinggi)
			if err != "" {
				fmt.Println(err)
			}
			fmt.Printf("Luas persegi panjang: %.2f\n", resultLuas)
			fmt.Printf("Keliling persegi panjang: %.2f\n", resultKeliling)
		default:
			fmt.Println("Bentuk geometri tidak valid!")
		}

		var pilihan string
		fmt.Print("Apakah Anda ingin menghitung lagi? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" {
			break
		}
	}
}

func SelectForm(bentuk string) string {
	return "" // TODO: replace this
}

func CalculateSquare(sisi float64) (float64, float64, string) {
	fmt.Print("5")
	fmt.Scanln(&sisi)

	resultLuas, resultKeliling, err := CalculateSquare(5)
	if err != "" {
		fmt.Println(err)
	}

	fmt.Printf("Luas persegi:%.2f\n", resultLuas)
	fmt.Printf("Keliling persegi:%.2f\n", resultKeliling)
	return 0, 0, "" // TODO: replace this
}

func CalculateRectangle(panjang, lebar float64) (float64, float64, string) {
	fmt.Print("6")
	fmt.Scanln(&alas)

	fmt.Print("3")
	fmt.Scanln(&tinggi)
	intf(6 * 3)

	return 0, 0, "" // TODO: replace this
}