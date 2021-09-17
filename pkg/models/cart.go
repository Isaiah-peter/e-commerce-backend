package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	dbc *gorm.DB
)

type Cart struct {
	gorm.Model
	UserId int64 `json:"user_id" binding:"required"`
	Product []ProductQty
}

type ProductQty struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64  `json:"quantity" binding:"required"`
}

func init() {
	config.Connect()
	dbc = config.GetDB()
	dbc.AutoMigrate(&Cart{}, &ProductQty{})
}
