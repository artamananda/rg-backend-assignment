package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	fileBytes, err := ioutil.ReadFile("assets/images/img-avatar.png")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Server Internal Error"})
		return
	}
	w.WriteHeader(http.StatusOK)

	api.dashboardView(w, r)
}
