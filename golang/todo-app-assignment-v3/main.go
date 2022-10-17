package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	e := model.ErrorResponse{}
	success := model.SuccessResponse{}
	code := 200
	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := decoder.Decode(&payload); err != nil {
		e.Error = "Internal Server Error"
		code = 400
	} else if payload.Username == "" || payload.Password == "" {
		e.Error = "Username or Password empty"
		code = 400
	} else if _, ok := db.Users[payload.Username]; ok {
		e.Error = "Username already exist"
		code = 409
	} else {
		success.Username = payload.Username
		success.Message = "Register Success"
		db.Users[payload.Username] = payload.Password
	}

	if e.Error != "" {
		jsonInBytes, err := json.Marshal(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	} else {
		jsonInBytes, err := json.Marshal(success)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	e := model.ErrorResponse{}
	success := model.SuccessResponse{}
	code := 200
	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := decoder.Decode(&payload); err != nil {
		e.Error = "Internal Server Error"
		code = 400
	} else if payload.Username == "" || payload.Password == "" {
		e.Error = "Username or Password empty"
		code = 400
	} else if _, ok := db.Users[payload.Username]; !ok {
		e.Error = "Wrong User or Password!"
		code = 401
	} else {
		success.Username = payload.Username
		success.Message = "Login Success"
		cookieName := "session_token"
		c := &http.Cookie{}
		if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
			c = storedCookie
		}
		if c.Value == "" {
			c = &http.Cookie{}
			c.Name = cookieName
			c.Value = uuid.NewString()
			c.Expires = time.Now().Add(5 * time.Hour)
			http.SetCookie(w, c)
		}

		db.Sessions[c.Name] = model.Session{
			Username: payload.Username,
			Expiry:   c.Expires,
		}
	}

	if e.Error != "" {
		jsonInBytes, err := json.Marshal(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	} else {
		jsonInBytes, err := json.Marshal(success)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	}
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	e := model.ErrorResponse{}
	success := model.SuccessResponse{}
	code := 200
	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Id   string `json:"id"`
		Task string `json:"task"`
		Done bool   `json:"done"`
	}{}

	if err := decoder.Decode(&payload); err != nil {
		e.Error = "Internal Server Error"
		code = 400
	} else {
		success.Username = db.Sessions["session_token"].Username
		success.Message = "Task Create a todo app program added!"
		res := model.Todo{}
		res.Id = uuid.NewString()
		res.Task = payload.Task
		res.Done = payload.Done
		db.Task[success.Username] = append(db.Task[success.Username], res)
	}

	if e.Error != "" {
		jsonInBytes, err := json.Marshal(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	} else {
		jsonInBytes, err := json.Marshal(success)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	}
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	e := model.ErrorResponse{}
	code := 200

	if len(db.Task) == 0 {
		e.Error = "Todolist not found!"
		code = 404
		jsonInBytes, err := json.Marshal(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	} else if len(db.Task[db.Sessions["session_token"].Username]) == 0 {
		jsonInBytes, err := json.Marshal(model.Todo{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	} else {
		jsonInBytes, err := json.Marshal(db.Task[db.Sessions["session_token"].Username])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(jsonInBytes)
	}
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	success := model.SuccessResponse{}
	code := 200

	success.Username = db.Sessions["session_token"].Username
	success.Message = "Clear ToDo Success"
	db.Task[success.Username] = []model.Todo{}

	jsonInBytes, err := json.Marshal(success)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonInBytes)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	success := model.SuccessResponse{}
	code := 200
	success.Username = db.Sessions["session_token"].Username
	success.Message = "Logout Success"

	jsonInBytes, err := json.Marshal(success)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonInBytes)
	delete(db.Sessions, "session_token")
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

	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(middleware.Auth(http.HandlerFunc(ClearToDo))))

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
