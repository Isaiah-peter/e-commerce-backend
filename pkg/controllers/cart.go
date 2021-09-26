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
	NewCart models.Cart
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	token:=utils.UseToken(r)
	cart := &models.Cart{}
    utils.ParseBody(r, cart)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	cart.UserId = verifiedID
	u:= cart.CreateCart()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "publication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	cart := &models.Cart{}
	utils.ParseBody(r, cart)
	vars := mux.Vars(r)
	productId := vars["id"]

	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		panic(err)
	}
	cartDetail, db := models.GetCartById(id)

	if cart.Product != nil {
		cartDetail.Product = cart.Product
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	db.Save(&cartDetail)
}


func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cart := &models.Cart{}
	utils.UseToken(r)
	vars := mux.Vars(r)
	productId := vars["id"]
	db.Where("ID=?", productId).Preload("Product").Delete(&cart)
	res, _ := json.Marshal("you have successfully delete this cart")
	w.Header().Set("Content-Type", "publication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCartById(w http.ResponseWriter, r *http.Request) {
	cart := &models.Cart{}
	token:= utils.UseToken(r)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	u:=db.Where("user_id=?", verifiedID).Preload("Product").Find(&cart).Value
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "publication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	var cart []models.Cart
	if token["IsAdmin"] == true {
		u := db.Preload("Product").Find(&cart).Value
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}