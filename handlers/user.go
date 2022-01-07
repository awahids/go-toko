package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/awahids/tokoDistributor/models"
	"gorm.io/gorm"
)

var db *gorm.DB

// Result is an array of user
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func AddUser(w http.ResponseWriter, router *http.Request) {
	payloads, _ := ioutil.ReadAll(router.Body)

	var user models.User
	json.Unmarshal(payloads, &user)

	db.Create(&user)

	res := Result{Code: 200, Data: user, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}