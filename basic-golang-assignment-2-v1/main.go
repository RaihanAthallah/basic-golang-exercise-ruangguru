package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	//students, _ := ConstructToSlice()
	message := ""
	if id != "" && name != "" {
		if len(id) == 5 {
			for _, student := range Students {
				splited := strings.Split(student, "_")
				fmt.Println(student)
				if id == splited[0] && name == splited[1] {
					name := splited[1]
					//program := splited[2]
					message = "Login berhasil: " + name
					break
				} else {
					message = "Login gagal: data mahasiswa tidak ditemukan"
					continue
				}
			}
		} else if len(id) != 5 {
			message = "ID must be 5 characters long!"
		}
	} else if len(id) == 0 || len(name) == 0 {
		message = "ID or Name is undefined!"
	}

	return message // TODO: replace this
}
func Register(id string, name string, major string) string {
	message := ""
	check := false
	if (id != "" && name != "") && major != "" {
		if len(id) == 5 {
			for _, student := range Students {
				splited := strings.Split(student, "_")
				if id == splited[0] {
					message = "Registrasi gagal: id sudah digunakan"
					check = true
					break
				}
			}
			if check == false {
				newStudents := id + "_" + name + "_" + major
				Students = append(Students, newStudents)
				message = "Registrasi berhasil: " + name + " (" + major + ")"
			}
		} else if len(id) != 5 {
			message = "ID must be 5 characters long!"
		}
	} else if len(id) == 0 || len(name) == 0 || len(major) == 0 {
		message = "ID, Name or Major is undefined!"
	}

	return message // TODO: replace this
}

func GetStudyProgram(code string) string {
	message := ""
	if code != "" {
		for key, major := range StudentStudyPrograms {
			//splited := strings.Split(major, "_")
			//fmt.Println(key)
			//fmt.Println(major)
			if code == key {
				message = major
				break
			} else {
				message = "Kode program studi tidak ditemukan"
				continue
			}
		}
	} else if len(code) == 0 {
		message = "Code is undefined!"
	}
	return message // TODO: replace this
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	message := ""
	if nama != "" {
		for i, _ := range Students {
			//splited := strings.Split(student, "_")
			if strings.Contains(Students[i], nama) {
				//print("Nama Ditemukan")
				//student = strings.Replace(student, splited[2], programStudi, 1)
				//fmt.Println(student)
				fn(programStudi, &Students[i])

				message = "Program studi mahasiswa berhasil diubah."
				//fmt.Println(Students[i])

				break
			} else {
				//print("Nama Tidak Ditemukan")
				message = "Mahasiswa tidak ditemukan."
				continue
			}
		}
	}

	return message // TODO: replace this
}

func UpdateStudyProgram(programStudi string, students *string) {
	splited := strings.Split(*students, "_")
	*students = strings.Replace(*students, splited[2], programStudi, 1)
	//updatedStudents := strings.Replace(*students, splited[2], programStudi, 1)
	//fmt.Println(updatedStudent)
	// TODO: answer here
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

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
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
