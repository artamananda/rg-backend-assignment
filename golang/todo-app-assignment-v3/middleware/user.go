package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		storedCookie, _ := r.Cookie("session_token")
		if isExpired(db.Sessions["session_token"]) || storedCookie == nil {
			e := model.ErrorResponse{}
			code := 401
			e.Error = "http: named cookie not present"
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
