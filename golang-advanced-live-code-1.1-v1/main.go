package main

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (api *API) ResetData(file string) error {
	jsonData, err := json.Marshal([]interface{}{})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/"+file, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) ChangeData(questions []model.Question) error {
	data, err := json.MarshalIndent(questions, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/questions.json", data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) ReadData() ([]model.Question, error) {
	file, err := ioutil.ReadFile("data/questions.json")
	if err != nil {
		return []model.Question{}, err
	}

	var questions []model.Question
	err = json.Unmarshal(file, &questions)
	if err != nil {
		return []model.Question{}, err
	}

	return questions, nil
}

func (api *API) AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
	dataQuestion := model.Question{}
	err := json.NewDecoder(r.Body).Decode(&dataQuestion)
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error: "Bad Request",
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse.Error))
		return
	}
	allData, err := api.ReadData()
	allData = append(allData, dataQuestion)
	_ = api.ResetData("questions.json")
	err = api.ChangeData(allData)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Question added!"))
}
func (api *API) GetAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	allData, err := api.ReadData()
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error: "Bad Request",
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse.Error))
		return
	}
	data, err := json.MarshalIndent(allData, "", "  ")
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error: "Bad Request",
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse.Error))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (api *API) GetQuestionByIDHandler(w http.ResponseWriter, r *http.Request) {
	allData, err := api.ReadData()
	if err != nil {
		errorResponse := model.ErrorResponse{
			Error: "Bad Request",
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorResponse.Error))
		return
	}
	id := r.URL.Query().Get("id")
	isFound := false
	for _, data := range allData {
		if data.ID == id {
			data, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				errorResponse := model.ErrorResponse{
					Error: "Bad Request",
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorResponse.Error))
				return
			}
			isFound = true
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
	}
	if !isFound {
		errorResponse := model.ErrorResponse{
			Error: "Question not found!",
		}
		errorResponseData, _ := json.MarshalIndent(errorResponse, "", "  ")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorResponseData)
		return
	}

}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/question/add", http.HandlerFunc(api.AddQuestionHandler))
	mux.Handle("/question/get-all", http.HandlerFunc(api.GetAllQuestionsHandler))
	mux.Handle("/question/get-by-id", http.HandlerFunc(api.GetQuestionByIDHandler))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
