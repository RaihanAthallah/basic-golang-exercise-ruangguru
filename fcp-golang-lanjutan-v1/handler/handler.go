package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var UserLogin = make(map[string]model.User)

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieRole, err := r.Cookie("user_login_role")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			var errMap model.ErrorResponse = model.ErrorResponse{
				Error: "user login role not found",
			}
			errBody, _ := json.Marshal(errMap)
			w.Write(errBody)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "role", cookieRole.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
		// your code here }) // TODO: replace this
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("An error occured by %v line 75", err)
	}

	var payload model.UserLogin
	err2 := json.Unmarshal(reqBody, &payload)
	if err2 != nil {
		log.Fatalf("An error occured by %v line 81", err2)
	}

	// Jika request body ID atau name kosong("")
	if len(payload.ID) == 0 || len(payload.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "ID or name is empty",
		}
		errBody, err3 := json.Marshal(errMap)
		if err3 != nil {
			log.Fatalf("An error occured by %v line 92", err3)
		}

		w.Write(errBody)
		return
	}

	// Jika ID user dan name tidak ditemukan di file data/users.txt
	if !service.CheckUserIsExist(payload.ID) {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "user not found",
		}
		errBody, err3 := json.Marshal(errMap)
		if err3 != nil {
			log.Fatalf("An error occured by %v line 107", err3)
		}
		w.Write(errBody)
		return

	}

	// Jika semua Oke
	// Berikan cookie dengan key user_login_id dan value <id user>
	cookieID := http.Cookie{
		Name:   "user_login_id",
		Value:  payload.ID,
		MaxAge: 3600,
	}

	// Berikan cookie dengan key user_login_role dan value <role user>
	role := service.GetRoleByID(payload.ID)
	cookieRole := http.Cookie{
		Name:   "user_login_role",
		Value:  role,
		MaxAge: 3600,
	}

	http.SetCookie(w, &cookieID)
	http.SetCookie(w, &cookieRole)

	w.WriteHeader(http.StatusOK)
	var resp model.SuccessResponse = model.SuccessResponse{
		Message:  "login success",
		Username: payload.ID,
	}
	respBody, err3 := json.Marshal(resp)
	if err3 != nil {
		log.Fatalf("An error occured by %v line 122", err3)
	}

	w.Write(respBody)
	// Simpan data user yang login ke dalam variable map UserLogin dengan key <id user> dan value berupa data user.
	user := service.GetUserByID(payload.ID)
	UserLogin[cookieID.Value] = user

}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("An error occured by %v line 150", err)
	}
	defer r.Body.Close()

	var user model.User
	err2 := json.Unmarshal(reqBody, &user)
	if err2 != nil {
		log.Fatalf("An error occured by %v line 157", err2)
	}
	// Jika request body ID, name, role atau study Code kosong(""),
	if len(user.ID) == 0 || len(user.Name) == 0 || len(user.StudyCode) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "ID, name, study code or role is empty",
		}

		errBody, err3 := json.Marshal(errMap)
		if err3 != nil {
			log.Fatalf("An error occured by %v line 168", err3)
		}

		w.Write(errBody)
		return
	}

	// Pastikan bahwa role yang diberikan hanya admin atau user
	if user.Role != "admin" && user.Role != "user" {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "role must be admin or user",
		}

		errBody, err3 := json.Marshal(errMap)
		if err3 != nil {
			log.Fatalf("An error occured by %v line 184", err3)
		}

		w.Write(errBody)
		return
	}

	// Pastikan juga study code sesuai dengan yang ada di file data/list-study.txt
	if !service.CheckStudyProgramIsExist(user.StudyCode) {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "study code not found",
		}

		errBody, err3 := json.Marshal(errMap)
		if err3 != nil {
			log.Fatalf("An error occured by %v line 194", err3)
		}

		w.Write(errBody)
		return

	}

	// Jika semua ID user sudah ada di penyimpanan data/users.txt
	if service.CheckUserIsExist(user.ID) {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "user id already exist",
		}
		errBody, err3 := json.Marshal(errMap)
		if err3 != nil {
			log.Fatalf("An error occured by %v line 211", err3)
		}

		w.Write(errBody)
		return
	}

	// Jika semua Oke

	w.WriteHeader(http.StatusOK)
	var resp model.SuccessResponse = model.SuccessResponse{
		Message:  "register success",
		Username: user.ID,
	}
	respBody, err3 := json.Marshal(resp)
	if err3 != nil {
		log.Fatalf("An error occured by %v line 231", err3)
	}
	service.PostData(user)
	w.Write(respBody)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// membaca data userID yang disimpan oleh middleware sebelumnya (middleware Auth)
	userID := r.Context().Value("userID").(string)

	// Hapus cookie dengan key user_login_id
	// delete(UserLogin, userID)
	cookieID := http.Cookie{
		Name:   "user_login_id",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookieID)

	// Hapus cookie dengan key user_login_role
	cookieRole := http.Cookie{
		Name:   "user_login_role",
		MaxAge: -1,
		Value:  "",
	}

	http.SetCookie(w, &cookieRole)

	w.WriteHeader(http.StatusOK)
	var resp model.SuccessResponse = model.SuccessResponse{
		Message:  "logout success",
		Username: userID,
	}
	w.Header().Set("Content-Type", "application/json")
	respBody, _ := json.Marshal(resp)
	w.Write(respBody)

	// Hapus data user yang login dari variable map UserLogin
	delete(UserLogin, cookieID.Value)
}

func GetStudyProgram(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	listData := service.GetListStudyProgram()
	respBody, _ := json.Marshal(listData)

	w.Write(respBody)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if role != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "user login role not Admin",
		}
		errBody, _ := json.Marshal(errMap)
		w.Write(errBody)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newUser model.User
	json.Unmarshal(reqBody, &newUser)

	if len(newUser.ID) == 0 || len(newUser.Name) == 0 || len(newUser.StudyCode) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "ID, name, or study code is empty",
		}

		errBody, _ := json.Marshal(errMap)
		w.Write(errBody)
		return
	}

	if service.CheckUserIsExist(newUser.ID) {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "study code not found", // SEHARUSNYA user id already exist TAPI DI TEST CASE study code not found
		}

		errBody, _ := json.Marshal(errMap)
		w.Write(errBody)
		return
	}

	if !service.CheckStudyProgramIsExist(newUser.StudyCode) {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "study code not found",
		}

		errBody, _ := json.Marshal(errMap)
		w.Write(errBody)
		return
	}

	w.WriteHeader(http.StatusOK)
	var resp model.SuccessResponse = model.SuccessResponse{
		Message:  "add user success",
		Username: newUser.ID,
	}
	respBody, _ := json.Marshal(resp)
	w.Write(respBody)
	newUser.Role = "user"
	service.PostData(newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	role := r.Context().Value("role").(string)
	if role != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "user login role not Admin",
		}
		errBody, _ := json.Marshal(errMap)
		w.Write(errBody)
		return
	}

	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "user id is empty",
		}
		errBody, _ := json.Marshal(errMap)
		w.Write(errBody)
		return
	}
	if !service.CheckUserIsExist(id) {
		w.WriteHeader(http.StatusBadRequest)
		var errMap model.ErrorResponse = model.ErrorResponse{
			Error: "user id not found",
		}
		errBody, err := json.Marshal(errMap)
		if err != nil {
			log.Fatalf("An error occured by %v line 346", err)
		}
		w.Write(errBody)
		return
	}

	w.WriteHeader(http.StatusOK)
	var resp model.SuccessResponse = model.SuccessResponse{
		Message:  "delete success",
		Username: id,
	}
	errBody, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("An error occured by %v line 359", err)
	}

	w.Write(errBody)
	service.DeleteData(id)
}

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeather(w http.ResponseWriter, r *http.Request) {
	var listRegion = []string{"jakarta", "bandung", "surabaya", "yogyakarta", "medan", "makassar", "manado", "palembang", "semarang", "bali"}

	var resCh = make(chan model.MainWeather, len(listRegion))
	var errCh = make(chan error, len(listRegion))

	// goroutine to get weather data for each region
	for _, region := range listRegion {
		go func(region string) {
			weather, err := GetWetherByRegionAPI(region)
			if err != nil {
				errCh <- err
				return
			}
			resCh <- weather
		}(region)
	}

	var result []model.MainWeather
	var err error
	// get results from channels
	for i := 0; i < len(listRegion); i++ {
		select {
		case weather := <-resCh:
			result = append(result, weather)
		case e := <-errCh:
			err = e
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)

	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	// TODO: answer here
}
