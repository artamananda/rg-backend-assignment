package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			e := model.ErrorResponse{}
			e.Error = "Method is not allowed!"
			code := 405
			jsonInBytes, err := json.Marshal(e)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			w.Write(jsonInBytes)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			e := model.ErrorResponse{}
			e.Error = "Method is not allowed!"
			code := 405
			jsonInBytes, err := json.Marshal(e)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			w.Write(jsonInBytes)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			e := model.ErrorResponse{}
			e.Error = "Method is not allowed!"
			code := 405
			jsonInBytes, err := json.Marshal(e)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			w.Write(jsonInBytes)
			return
		}
		next.ServeHTTP(w, r)
	})
}
