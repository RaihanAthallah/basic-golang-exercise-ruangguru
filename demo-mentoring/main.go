package main

import (
	"html/template"
	"net/http"
)

func getHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		text := "Hello " + name
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(text))
	}
}

func getHelloHtml() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		text := "Hello " + name

		template, err := template.ParseFiles("./template/main.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}

		template.Execute(w, text)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(text))
	}

}

func main() {

	http.HandleFunc("/hello", getHello())
	http.HandleFunc("/hellohtml", getHelloHtml())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Main page"))
	})

	http.ListenAndServe(":8080", nil)
}
