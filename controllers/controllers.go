package controllers

import (
	"encoding/json"
	"fmt"
	"klever/api/database"
	"klever/api/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page atual")
}


func Details(w http.ResponseWriter, r *http.Request) {
	var d []models.Detail
	database.DB.Find(&d)
	json.NewEncoder(w).Encode(d)
}


func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction models.Transaction
	json.NewDecoder(r.Body).Decode(&newTransaction)
	database.DB.Create(&newTransaction)
	json.NewEncoder(w).Encode(newTransaction)
}


func ReceiveTransaction(w http.ResponseWriter, r *http.Request) {

}


func Balances(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Balances)
}
