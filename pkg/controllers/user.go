package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/models"
	utils "github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)



var (
	db = config.GetDB()
	NewUser models.User
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	utils.ParseBody(r, user)
	u := user.CreateUser()
	res, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	utils.ParseBody(r, user)
	u := FindOne(user.Email, user.Password)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(u)
}

func FindOne(email string, password string) map[string]interface{} {
	var user = &models.User{}
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expireAt := time.Now().Add(time.Minute * 50).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID:  int64(user.ID),
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	token := utils.UseToken(r)
	utils.ParseBody(r, user)
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}

	if verifiedID == id || token["IsAdmin"] == true {
		userDetail, dr := models.GetUserById(id)

		if user.Password != "" {
			hashpassword, err := utils.HashPassword(user.Password)
			if err != nil {
				panic(err)
			}
			userDetail.Password = hashpassword
		}
		if user.UserName != "" {
			userDetail.UserName = user.UserName
		}
		if user.Email != "" {
			userDetail.Email = user.Email
		}
		if user.IsAdmin != false {
			userDetail.IsAdmin = user.IsAdmin
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		dr.Save(&userDetail)
	}

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	vars := mux.Vars(r)
	userId := vars["id"]
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}
	if verifiedID == id || token["IsAdmin"] == true {
		userDetail, _ := models.GetUserById(id)
		res, _ := json.Marshal(userDetail)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	userId := vars["id"]

	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}

	if verifiedID == id || token["IsAdmin"] == true {
		user := models.DeleteUser(id)
		res, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	token := utils.UseToken(r)
	new := r.URL.Query()["new"]
	var user []models.User
	u := db.Find(&user).Value
	u_five := db.Find(&user).Order("user_name ASC").Limit(5).Value
	if token["IsAdmin"] == true {
		if len(new) != 0 {
			res, _ := json.Marshal(u_five)
			w.Header().Set("Content-Type", "publication/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}else {
			res, _ := json.Marshal(u)
			w.Header().Set("Content-Type", "publication/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
	}else {
		res,_ := json.Marshal( "you are not an admin")
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

}


func GetUserUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query()["username"]
	token := utils.UseToken(r)
	var user []models.User
	u := db.Where("user_name=?", username).Find(&user).Value
	if token["IsAdmin"] == true {
		res, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}else {
		res,_ := json.Marshal( "you are not an admin")
		w.Header().Set("Content-Type", "publication/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
	}

}