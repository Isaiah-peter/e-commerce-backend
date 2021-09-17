package models

import (
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var (
	dbu *gorm.DB
)

type User struct {
	gorm.Model
	UserName string `json:"user_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" `
	IsAdmin bool `json:"is_admin"`
}

type Token struct {
	UserID  int64
	IsAdmin bool
    jwt.StandardClaims
}

func init() {
	config.Connect()
	dbu = config.GetDB()
	dbu.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	u.Password = hashPassword
    dbc.NewRecord(u)
    dbc.Create(u)
    return u
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := dbu.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	dbu.Where("ID=?", Id).Delete(user)
	return user
}