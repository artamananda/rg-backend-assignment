package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{}
	code := 200
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		if err := decoder.Decode(&payload); err != nil {
			message["error"] = "Internal Server Error"
			code = 400

		} else if payload.Username == "" || payload.Password == "" {
			message["error"] = "Username or Password empty"
			code = 400
		} else if _, ok := db.Users[payload.Username]; ok {
			message["error"] = "Username already exist"
			code = 409
		} else {
			message["username"] = payload.Username
			message["message"] = "Register Success"
		}

	} else {
		message["error"] = "Method is not allowed!"
		code = 405
	}

	jsonInBytes, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonInBytes)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Logout(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	// TODO: answer here
	fmt.Print(username)
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	// TODO: answer here

	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

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
