package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var res string
		t := time.Now()
		res += fmt.Sprint(t.Weekday()) + ", "
		res += fmt.Sprint(t.Day()) + " "
		res += fmt.Sprint(t.Month()) + " "
		res += fmt.Sprint(t.Year())

		w.WriteHeader(200)
		w.Write([]byte(res))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		if r.URL.Query().Has("name") {
			name := r.URL.Query().Get("name")
			w.Write([]byte("Hello, " + name + "!"))
		} else {
			w.Write([]byte("Hello there"))
		}
	}) // TODO: replace this
}

func main() {
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())
	http.ListenAndServe("localhost:8080", nil)
}
