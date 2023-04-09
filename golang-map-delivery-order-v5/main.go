package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)
	for _, v := range data {
		sliced := strings.Split(v, ":")
		harga, _ := strconv.ParseFloat(sliced[2], 32)
		nama := sliced[0] + "-" + sliced[1]
		admin := 0.0
		if sliced[3] == "JKT" && day != "minggu" {
			if day == "senin" || day == "rabu" || day == "jumat" {
				admin = harga * 0.1
				harga = harga + admin
			} else if day == "selasa" || day == "kamis" || day == "sabtu" {
				admin = harga * 0.05
				harga = harga + admin
			}

			result[nama] = float32(harga)
		} else if sliced[3] == "BDG" && (day == "rabu" || day == "kamis" || day == "sabtu") {
			if day == "senin" || day == "rabu" || day == "jumat" {
				admin = harga * 0.1
				harga = harga + admin
			} else if day == "selasa" || day == "kamis" || day == "sabtu" {
				admin = harga * 0.05
				harga = harga + admin
			}
			result[nama] = float32(harga)
		} else if sliced[3] == "BKS" && (day == "selasa" || day == "kamis" || day == "jumat") {
			if day == "senin" || day == "rabu" || day == "jumat" {
				admin = harga * 0.1
				harga = harga + admin
			} else if day == "selasa" || day == "kamis" || day == "sabtu" {
				admin = harga * 0.05
				harga = harga + admin
			}
			result[nama] = float32(harga)
		} else if sliced[3] == "DPK" && (day == "senin" || day == "selasa") {
			if day == "senin" || day == "rabu" || day == "jumat" {
				admin = harga * 0.1
				harga = harga + admin
			} else if day == "selasa" || day == "kamis" || day == "sabtu" {
				admin = harga * 0.05
				harga = harga + admin
			}
			result[nama] = float32(harga)
		}
	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
