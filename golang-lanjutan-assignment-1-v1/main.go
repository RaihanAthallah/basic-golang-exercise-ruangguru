package main

import (
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// type StudyData struct {
// 	Code string `json:"code"`
// 	Name string `json:"name"`
// }

// type ErrorResponse struct {
// 	Error string `json:"error"`
// }

func isCheckUserExist(id string) bool {
	file, err := os.OpenFile("data/users.txt", os.O_RDONLY, 0644)
	check := false
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splited := strings.Split(text, "_")
		if splited[0] == id {
			check = true
		}
	}

	defer file.Close()
	return check
}

func isCheckStudyExist(code string) bool {
	file, err := os.OpenFile("data/list-study.txt", os.O_RDONLY, 0644)
	fmt.Println(code)
	check := true
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splited := strings.Split(text, "_")
		if splited[0] == code {
			fmt.Println(code)
			fmt.Println(splited[0])
			check = false
		}
	}

	defer file.Close()
	return check
}

func GetStudyProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studyData []model.StudyData

		if r.Method != "GET" {
			errorMessage := model.ErrorResponse{Error: "Method is not allowed!"}
			errorMessageJson, _ := json.Marshal(errorMessage)

			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(errorMessageJson)
		} else {

			file, err := os.OpenFile("data/list-study.txt", os.O_RDONLY, 0644)
			if err != nil {
				w.Write([]byte("Error: " + err.Error()))

			}
			fileScanner := bufio.NewScanner(file)
			for fileScanner.Scan() {
				text := fileScanner.Text()
				splited := strings.Split(text, "_")
				studyData = append(studyData, model.StudyData{
					Code: splited[0],
					Name: splited[1],
				})
			}

			defer file.Close()

			studyDataJson, _ := json.Marshal(studyData)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(studyDataJson)
		}
	}
}

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			errorMessage := model.ErrorResponse{Error: "Method is not allowed!"}
			errorMessageJson, _ := json.Marshal(errorMessage)

			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(errorMessageJson)
		} else {
			var user model.User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusOK)
				return
			}
			// fmt.Println(user.StudyCode)
			if user.ID == "" || user.Name == "" || user.StudyCode == "" {
				errorMessage := model.ErrorResponse{Error: "ID, name, or study code is empty"}
				errorMessageJson, _ := json.Marshal(errorMessage)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorMessageJson)
				return

			} else if isCheckStudyExist(user.StudyCode) {
				fmt.Println("masuk")
				errorMessage := model.ErrorResponse{Error: "study code not found"}
				errorMessageJson, _ := json.Marshal(errorMessage)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorMessageJson)
				return

			} else if isCheckUserExist(user.ID) {
				errorMessage := model.ErrorResponse{Error: "User already exists"}
				errorMessageJson, _ := json.Marshal(errorMessage)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorMessageJson)
				return

			} else {
				file, err := os.OpenFile("data/users.txt", os.O_APPEND|os.O_WRONLY, 0644)

				if err != nil {
					w.Write([]byte("Error: " + err.Error()))
				}

				_, err = file.WriteString(user.ID + "_" + user.Name + "_" + user.StudyCode + "\n")

				defer file.Close()

				if err != nil {
					w.Write([]byte("Error: " + err.Error()))
				}

				successMessage := model.SuccessResponse{Username: user.Name, Message: "add user success"}
				successMessageJson, _ := json.Marshal(successMessage)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(successMessageJson))

			}

		}
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			errorMessage := model.ErrorResponse{Error: "Method is not allowed!"}
			errorMessageJson, _ := json.Marshal(errorMessage)

			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(errorMessageJson)
		} else {
			id := r.URL.Query().Get("id")

			if id == "" {
				errorMessage := model.ErrorResponse{Error: "user id is empty"}
				errorMessageJson, _ := json.Marshal(errorMessage)

				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorMessageJson)
				return
			} else if !isCheckUserExist(id) {
				errorMessage := model.ErrorResponse{Error: "user id not found"}
				errorMessageJson, _ := json.Marshal(errorMessage)

				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorMessageJson)
				return
			} else {

				file, err := os.OpenFile("data/users.txt", os.O_RDONLY, 0644)
				if err != nil {
					w.Write([]byte("Error: " + err.Error()))

				}
				fileScanner := bufio.NewScanner(file)

				var username string

				for fileScanner.Scan() {
					text := fileScanner.Text()
					splited := strings.Split(text, "_")
					if splited[0] == id {
						username = splited[1]
					}
				}

				successMessage := model.SuccessResponse{Username: username, Message: "delete success"}
				successMessageJson, _ := json.Marshal(successMessage)

				w.WriteHeader(http.StatusOK)
				w.Write(successMessageJson)

				// defer file.Close()
			}
		}
	}
}

func main() {
	http.HandleFunc("/study-program", GetStudyProgram())
	http.HandleFunc("/user/add", AddUser())
	http.HandleFunc("/user/delete", DeleteUser())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
