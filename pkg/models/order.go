package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbo *gorm.DB
)


type Order struct {
	gorm.Model
	UserId int64 `json:"user_id" binding:"required"`
	Product []ProductQty `json:"product"`
	Amount  int64 `json:"Amount" binding:"required"`
	Address Address `json:"address"`
	Status string `json:"status" default:"pending"`
}

type Address struct {
	 gorm.Model
	 Addresses string `json:"addresses" binding:"required"`
}

func init() {
	config.Connect()
	dbo = config.GetDB()
	dbc.AutoMigrate(&Order{}, &Address{})
}


