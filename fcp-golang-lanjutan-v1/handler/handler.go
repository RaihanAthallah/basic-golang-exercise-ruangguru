
package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var UserLogin = make(map[string]model.User)

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

func getStudyPrograms() []model.StudyData {
	file, err := os.OpenFile("data/list-study.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	studyData := []model.StudyData{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splited := strings.Split(text, "_")
		studyData = append(studyData, model.StudyData{Code: splited[0], Name: splited[1]})
	}
	return studyData
}
func getUsers() []model.User {
	file, err := os.OpenFile("data/users.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	usersData := []model.User{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splited := strings.Split(text, "_")
		usersData = append(usersData, model.User{ID: splited[0], Name: splited[1], Role: splited[2], StudyCode: splited[3]})
	}
	return usersData
}

func getUserRole(id string) string {
	file, err := os.OpenFile("data/users.txt", os.O_RDONLY, 0644)
	role := ""
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		splited := strings.Split(text, "_")
		if splited[0] == id {
			role = splited[3]
			break
		}
	}
	return role
}

// DESC: func Auth is a middleware to check user login id, only user that already login can pass this middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_login_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if _, ok := UserLogin[c.Value]; !ok || c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", c.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DESC: func AuthAdmin is a middleware to check user login role, only admin can pass this middleware
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // your code here }) // TODO: replace this
		c, err := r.Cookie("user_login_role")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
		if c.Value != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "admin login id not found"})
			return
		}
		// ctx := r.Context()
		// ctx = context.WithValue(ctx, "userID", c.Value)
		next.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	// Check HTTP method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method is not allowed!"})
		return
	}

	// Parse request body
	var reqBody struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Check if ID or Name is empty
	if reqBody.ID == "" || reqBody.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID or name is empty"})
		return
	}

	// Get user data
	// userData := isCheckUserExist(reqBody.ID)
	if !isCheckUserExist(reqBody.ID) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	// Create cookies
	cookieID := &http.Cookie{Name: "user_login_id", Value: reqBody.ID}
	cookieRole := &http.Cookie{Name: "user_login_role", Value: getUserRole(reqBody.ID)}
	userData := model.User{ID: reqBody.ID, Name: reqBody.Name, Role: getUserRole(reqBody.ID)}
	// Set cookies in response
	http.SetCookie(w, cookieID)
	http.SetCookie(w, cookieRole)

	// Add user to UserLogin map
	UserLogin[reqBody.ID] = userData

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"username": reqBody.ID, "message": "login success"})
}

func Register(w http.ResponseWriter, r *http.Request) {
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
		if user.ID == "" || user.Name == "" || user.StudyCode == "" || user.Role == "" {
			errorMessage := model.ErrorResponse{Error: "ID, name, study code or role is empty"}
			errorMessageJson, _ := json.Marshal(errorMessage)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorMessageJson)
			return

		} else if user.Role != "admin" && user.Role != "user" {
			errorMessage := model.ErrorResponse{Error: "role must be admin or user"}
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
			errorMessage := model.ErrorResponse{Error: "user id already exist"}
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

			_, err = file.WriteString(user.ID + "_" + user.Name + "_" + user.StudyCode + "_" + user.Role + "\n")

			defer file.Close()
			if err != nil {
				w.Write([]byte("Error: " + err.Error()))
			}

			successMessage := model.SuccessResponse{Username: user.Name, Message: "register success"}
			successMessageJson, _ := json.Marshal(successMessage)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(successMessageJson))

		}

	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	if r.Method != "POST" {
		errorMessage := model.ErrorResponse{Error: "Method is not allowed!"}
		errorMessageJson, _ := json.Marshal(errorMessage)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(errorMessageJson)
	} else {
		_, err := r.Cookie("user_login_id")
		if err != nil {
			if err == http.ErrNoCookie {
				errorMessage := model.ErrorResponse{Error: "user login id not found"}
				errorMessageJson, _ := json.Marshal(errorMessage)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(errorMessageJson)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cookie := http.Cookie{
			Name:   "user_login_id",
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{
			Name:   "user_login_role",
			Value:  "",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)

		successMessage := model.SuccessResponse{Username: userID, Message: "logout success"}
		successMessageJson, _ := json.Marshal(successMessage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(successMessageJson)

	}

	// TODO: answer here
}

func GetStudyProgram(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorMessage := model.ErrorResponse{Error: "Method is not allowed!"}
		errorMessageJson, _ := json.Marshal(errorMessage)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(errorMessageJson)
	} else {
		// _, err := r.Cookie("user_login_id")
		// if err != nil {
		// 	if err == http.ErrNoCookie {
		// 		errorMessage := model.ErrorResponse{Error: "user login id not found"}
		// 		errorMessageJson, _ := json.Marshal(errorMessage)
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		w.Write(errorMessageJson)
		// 		return
		// 	}
		// }
		studyPrograms := getStudyPrograms()
		studyProgramsJson, _ := json.Marshal(studyPrograms)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(studyProgramsJson)
	}
	// TODO: answer here
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	user := model.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.ID == "" || user.Name == "" || user.StudyCode == "" || user.Role == "" {
		errorMessage := model.ErrorResponse{Error: "ID, name, study code or role is empty"}
		errorMessageJson, _ := json.Marshal(errorMessage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessageJson)
		return

	} else if isCheckStudyExist(user.StudyCode) {
		// fmt.Println("masuk")
		errorMessage := model.ErrorResponse{Error: "study code not found"}
		errorMessageJson, _ := json.Marshal(errorMessage)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessageJson)

		return

	} else if isCheckUserExist(user.ID) {
		errorMessage := model.ErrorResponse{Error: "user id already exist"}
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

		_, err = file.WriteString(user.ID + "_" + user.Name + "_" + user.StudyCode + "_" + user.Role + "\n")

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

	// TODO: answer here
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	userID := r.URL.Query().Get("id")

	if userID == "" {
		errorMessage := model.ErrorResponse{Error: "user id is empty"}
		errorMessageJson, _ := json.Marshal(errorMessage)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessageJson)
		return
	}

	// var updatedUsers []model.User
	var isFound bool

	for _, user := range users {
		if user.ID == userID {
			// updatedUsers = append(updatedUsers, user)
			file, err := os.OpenFile("data/users.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				w.Write([]byte("Error: " + err.Error()))
			}

			_, err = file.WriteString(user.ID + "_" + user.Name + "_" + user.StudyCode + "_" + user.Role + "\n")

			defer file.Close()
			if err != nil {
				w.Write([]byte("Error: " + err.Error()))
			}
		} else {
			isFound = true
		}
	}

	if !isFound {
		errorMessage := model.ErrorResponse{Error: "user id not found"}
		errorMessageJson, _ := json.Marshal(errorMessage)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessageJson)
	}

	successMessage := model.SuccessResponse{Username: userID, Message: "delete success"}
	successMessageJson, _ := json.Marshal(successMessage)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(successMessageJson))
	// TODO: answer here
}

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeather(w http.ResponseWriter, r *http.Request) {
	var listRegion = []string{"jakarta", "bandung", "surabaya", "yogyakarta", "medan", "makassar", "manado", "palembang", "semarang", "bali"}

	var resCh = make(chan model.MainWeather, len(listRegion))
	var errCh = make(chan error, len(listRegion))

	for _, region := range listRegion {
		go func(resCh chan model.MainWeather, region string) {
			weather, err := GetWetherByRegionAPI(region)
			if err != nil {
				errCh <- err
			} else {
				resCh <- weather
			}

		}(resCh, region)
	}

	var weathers []model.MainWeather

	for i := 0; i < len(listRegion); i++ {
		select {
		case weather := <-resCh:
			weathers = append(weathers, weather)
		case err := <-errCh:
			errorMessage := model.ErrorResponse{Error: err.Error()}
			errorMessageJson, _ := json.Marshal(errorMessage)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorMessageJson)
			return
		}
	}

	weathersJson, _ := json.Marshal(weathers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(weathersJson)

	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	// TODO: answer here
}
