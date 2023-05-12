package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}
	data := []Animechan{}
	req, err := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes/anime?title=naruto", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// responseData
	// data = append(data, Animechan{
	// 	Anime:     "Naruto",
	// 	Character: "Naruto Uzumaki",

	// })

	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, err
	}

	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method GET:
	return data, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	data := Postman{}
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	// resp, err := http.NewRequest("POST", "https://postman-echo.com/post", requestBody)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	// Hit API https://postman-echo.com/post with method POST:
	return data, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
