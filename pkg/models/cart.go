package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)


type Cart struct {
	gorm.Model
	UserId     int64 `json:"user_id" binding:"required"`
	ProductID  int64 `json:"product_id" binding:"required"`
	Product    []ProductQty
	Color      string `json:"color"`
	Size       string `json:"size"`
	TotalPrice string `json:"total_price"`
	Quantity   int64  `json:"quantity" binding:"required"`
}

type ProductQty struct {
	gorm.Model
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required"`
	CartID    int64 `json:"cart_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Cart{}, &ProductQty{})
}

func (c *Cart) CreateCart() *Cart {
	db.NewRecord(c)
	db.Create(c)
	return c
}

func GetCartById(Id int64) (*Cart, *gorm.DB) {
	var getUser Cart
	db := db.Where("ID=?", Id).Preload("Product").Find(&getUser)
	return &getUser, db
}
