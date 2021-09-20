package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

var (
	NewProduct models.Product
	NewCat models.Category
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	product := &models.Product{}
	utils.ParseBody(r, product)
	if token["IsAdmin"] == true {
		u := product.CreateProduct()

		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}else {
		res, _ := json.Marshal("you are not a seller contact an admin to make you a seller")
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	category := &models.Category{}
	utils.ParseBody(r, category)
	if token["IsAdmin"] == true {
		u := category.CreateCategory()
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}else {
		res, _ := json.Marshal("you are not a seller contact an admin to make you a seller")
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product = &models.Product{}
	token := utils.UseToken(r)
	utils.ParseBody(r, product)
	vars := mux.Vars(r)
	productId := vars["id"]

	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		panic(err)
	}
	if token["IsAdmin"] == true {
		productDetail, dr := models.GetProductById(id)

		if product.Categories != nil {
			productDetail.Categories = product.Categories
		}

		if product.Color != "" {
			productDetail.Color = product.Color
		}
		if product.Price <= 0 {
			productDetail.Price = product.Price
		}

		if product.ImageUrl != "" {
			productDetail.ImageUrl = product.ImageUrl
		}
		if product.Desc != "" {
			productDetail.Desc = product.Desc
		}
		if product.Title != "" {
			productDetail.Title = product.Title
		}
		if product.Size != "" {
			productDetail.Size = product.Size
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		dr.Save(&productDetail)
	}else {
		res, _ := json.Marshal("you are not an admin ")
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	var product []models.Product
	utils.ParseBody(r, product)
	vars := mux.Vars(r)
	productId := vars["id"]

	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		panic(err)
	}
	if token["IsAdmin"] == true {
		u := db.Where("ID=?",id).Delete(&product)
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}else {
		res, _ := json.Marshal("you are not an admin ")
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}
		productDetail, _ := models.GetProductById(id)
		res, _ := json.Marshal(productDetail)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	utils.UseToken(r)
	var product []models.Product
	var new = r.URL.Query()["new"]
	var cat []models.Category
	var category = r.URL.Query()["categories"]
	var id []string
	if len(new) != 0 {
		productDetail := db.Find(&product).Order("`product`.`id` DESC").Limit(1)
		res, _ := json.Marshal(productDetail)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	if len(new) == 0 && len(category) == 0 {
		productDetail := db.Find(&product).Value
		res, _ := json.Marshal(productDetail)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	if len(category) != 0 {
		fmt.Println("category",category[0])
		db.Where("name=?",category[0]).Find(&cat).Pluck("product_id",id)
		fmt.Println(strings.Join(id, ","))
		u := db.Where("name=?", category[0]).Find(&cat)
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}


