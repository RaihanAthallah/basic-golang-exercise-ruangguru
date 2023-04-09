package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func ConstructToSlice() ([]string, []string) {
	studentSlice := strings.Split(Students, ",")
	programSlice := strings.Split(StudentStudyPrograms, ",")
	//fmt.Println(studentSlice)
	//fmt.Println(len(studentSlice))
	//fmt.Println(programSlice)
	//fmt.Println(len(programSlice))

	return studentSlice, programSlice
}

func Login(id string, name string) string {
	//check := false
	students, _ := ConstructToSlice()
	//for check == false {
	if id != "" && name != "" {
		if len(id) == 5 {
			for _, student := range students {
				splited := strings.Split(student, "_")
				if id == splited[0] && name == splited[1] {
					//fmt.Println(student)
					name := splited[1]
					program := splited[2]
					//fmt.Println(program)
					//fmt.Println(name)
					fmt.Printf("Login berhasil: %s (%s)", name, program)

					break
				} else {
					fmt.Printf("Login gagal: data mahasiswa tidak ditemukan")
					break
				}

			}
		} else if len(id) != 5 {
			fmt.Printf("ID must be 5 characters long!")
		}
	} else if id == "" && name == "" {
		fmt.Printf("ID or Name is undefined!")
	}

	//}
	//check := strings.HasPrefix()
	return "" // TODO: replace this
}

func Register(id string, name string, major string) string {
	return "" // TODO: replace this
}

func GetStudyProgram(code string) string {
	return "" // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")
	ConstructToSlice()
	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)
		ConstructToSlice()
		//fmt.Println(len(Students))

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
