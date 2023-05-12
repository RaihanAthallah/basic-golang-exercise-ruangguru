package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		year := now.Year()
		month := now.Month()
		day := now.Day()
		weekday := now.Weekday()
		result := fmt.Sprintf("%s, %d %s %d", weekday, day, month.String(), year)
		w.Write([]byte(result))
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("Hello there"))
		} else {
			w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())

	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
