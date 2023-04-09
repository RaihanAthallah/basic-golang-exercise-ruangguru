package main

import (
	"fmt"
)

func main() {
	for {
		var panjang float64
		var dari, ke string

		fmt.Println("=== Kalkulator Konversi Satuan Panjang ===")

		fmt.Print("Masukkan panjang: ")
		fmt.Scanln(&panjang)

		fmt.Print("Masukkan satuan dari (m/cm/ft/in): ")
		fmt.Scanln(&dari)

		fmt.Print("Masukkan satuan ke (m/cm/ft/in): ")
		fmt.Scanln(&ke)

		result := ConvertLength(panjang, dari, ke)
		fmt.Printf("%.2f %s = %.2f %s\n", panjang, dari, result, ke)

		var pilihan string
		fmt.Print("Apakah Anda ingin mengkonversi kembali? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" {
			break
		}
	}
}

func ConvertLength(panjang float64, dari, ke string) float64 {

	if panjang > 0 {
		if dari == "m" {
			if ke == "cm" {
				panjang *= 100
			} else if ke == "ft" {
				panjang *= 3.281

			} else if ke == "in" {
				panjang *= 39.37
			}
		} else if dari == "cm" {
			if ke == "m" {
				panjang /= 100
			} else if ke == "ft" {
				panjang /= 30.48
			} else if ke == "in" {
				panjang /= 2.54
			}

		} else if dari == "ft" {
			if ke == "m" {
				panjang /= 3.281
			} else if ke == "cm" {
				panjang *= 30.48
			} else if ke == "in" {
				panjang *= 12
			}
		} else if dari == "in" {
			if ke == "m" {
				panjang /= 39.37
			} else if ke == "cm" {
				panjang *= 2.54
			} else if ke == "ft" {
				panjang /= 12
			}
		} else {
			return -1
		}
	} else {
		return 0
	}
	return panjang // TODO: replace this
}
