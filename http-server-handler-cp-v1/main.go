package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		year, month, date := time.Now().Date()
		day := time.Now().Weekday()
		dateResult := fmt.Sprintf("%s, %d %s %d", day, date, month.String(), year)

		writer.Write([]byte(dateResult))
	} // TODO: replace this
}

func main() {
	http.HandleFunc("/", GetHandler())
	http.ListenAndServe("localhost:8080", GetHandler())
}
