package controllers

import (
	"encoding/json"
	"net/http"
	"simple-rest-api/models"
	u "simple-rest-api/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) // decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	response := user.Create()
	u.Respond(w, response)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	response := models.Login(user.Email, user.Password)
	u.Respond(w, response)
}
