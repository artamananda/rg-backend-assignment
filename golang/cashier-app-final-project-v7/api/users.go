package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	// Read username and password request with FormValue.
	creds := model.Credentials{}
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}
	// Handle request if creds is empty send response code 400, and message "Username or Password empty"

	err := api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	// Read usernmae and password request with FormValue.
	creds := model.Credentials{}
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}
	// Handle request if creds is empty send response code 400, and message "Username or Password empty"

	err := api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Wrong User or Password!"})
		return
	}

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.
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

	session := model.Session{}
	session.Token = c.Value
	session.Username = creds.Username
	session.Expiry = c.Expires
	err = api.sessionsRepo.AddSessions(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	//Read session_token and get Value:
	storedCookie, _ := r.Cookie("session_token")
	if storedCookie == nil {
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
	sessionToken := storedCookie.Value

	api.sessionsRepo.DeleteSessions(sessionToken)

	//Set Cookie name session_token value to empty and set expires time to Now:
	storedCookie.Name = "session_token"
	storedCookie.Value = ""
	storedCookie.Expires = time.Now()

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
