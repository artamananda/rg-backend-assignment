package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	var res string
	var result []byte
	t := time.Now()
	res += fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())

	for i := 0; i < len(res); i++ {
		result = append(result, byte(res[i]))
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write(result)
	}
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
