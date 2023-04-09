package main

import (
	"fmt"
)

func ViewStudents(students map[string][]interface{}) {
	fmt.Println(len(students))
	fmt.Println(students["Raihan"])
	fmt.Println("Name\tAddress\tPhone\tScore")
	for name, info := range students {
		address, _ := info[0].(string)
		phone, _ := info[1].(string)
		score, _ := info[2].(int)
		fmt.Printf("%s\t%s\t%s\t%d\n", name, address, phone, score)
	}
}

func AddStudent(students *map[string][]interface{}) func(string, string, string, int) {
	return func(name string, address string, phone string, score int) {
		(*students)[name] = []interface{}{address, phone, score}
	}
}

func RemoveStudent(students *map[string][]interface{}) func(string) {
	return func(name string) {
		delete(*students, name)
	} // TODO: replace this
}

func UpdateScore(students *map[string][]interface{}) func(string, int) {

	return func(name string, score int) {
		(*students)[name][2] = score
	}

}

func main() {
	students := make(map[string][]interface{})
	add := AddStudent(&students)
	remove := RemoveStudent(&students)
	update := UpdateScore(&students)

	for {
		var command string
		fmt.Print("Enter command (add, remove, update-score, view): ")
		fmt.Scan(&command)

		switch command {
		case "add":
			var name, address, phone string
			var score int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter address: ")
			fmt.Scan(&address)
			fmt.Print("Enter phone: ")
			fmt.Scan(&phone)
			fmt.Print("Enter score: ")
			fmt.Scan(&score)

			add(name, address, phone, score)
		case "remove":
			var name string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			remove(name)
		case "update-score":
			var score int
			var name string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter new score: ")
			fmt.Scan(&score)
			update(name, score)
			fmt.Println("Score updated:")
			ViewStudents(students)
		case "view":
			fmt.Println("Student data:")
			ViewStudents(students)
		default:
			fmt.Println("Invalid command. Available commands: add, remove, update-score, view")
		}
	}
}
