package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	// Get username context to struct model.Cart.
	username := fmt.Sprint(r.Context().Value("username"))

	r.ParseForm()

	// Check r.Form with key product, if not found then return response code 400 and message "Request Product Not Found".
	x := r.Form.Has("product")
	if !x {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Request Product Not Found"})
		return
	}

	var list []model.Product
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}

	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	cart := model.Cart{}
	cart.Name = username
	cart.Cart = list
	for _, val := range list {
		cart.TotalPrice += val.Total
	}

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	api.dashboardView(w, r)

}
