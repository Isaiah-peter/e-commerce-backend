package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	utils "github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	NewOrder models.Order
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	token:=utils.UseToken(r)
	cart := &models.Order{}
	utils.ParseBody(r, cart)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	cart.UserId = verifiedID
	cart.Status = "pending"
	u:= cart.CreateOrder()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "publication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	cart := &models.Order{}
	utils.ParseBody(r, cart)
	vars := mux.Vars(r)
	productId := vars["id"]

	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		panic(err)
	}
	cartDetail, db := models.GetOrderById(id)

	if cart.OrderQuantity != nil {
		cartDetail.OrderQuantity = cart.OrderQuantity
	}
	if cart.Amount != 0 {
		cartDetail.Amount = cart.Amount
	}

	if cart.Status != "" {
		cartDetail.Status = cart.Status
	}

	if cart.Address.Addresses != "" {
		cartDetail.Address.Addresses = cart.Address.Addresses
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	db.Save(&cartDetail)
}


func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	cart := &models.Order{}
	utils.UseToken(r)
	vars := mux.Vars(r)
	productId := vars["id"]
	db.Where("ID=?", productId).Preload("OrderQuantity").Delete(&cart)
	res, _ := json.Marshal("you have successfully delete this cart")
	w.Header().Set("Content-Type", "publication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	cart := &models.Order{}
	token:= utils.UseToken(r)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	u:=db.Where("user_id=?", verifiedID).Preload("OrderQuantity").Find(&cart).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "publication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	var cart []models.Order
	if token["IsAdmin"] == true {
		u := db.Preload("OrderQuantity").Find(&cart).Value
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}


