package models

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/config"
	"github.com/jinzhu/gorm"
)


type Order struct {
	gorm.Model
	UserId          int64 `json:"user_id" binding:"required"`
	OrderQuantity   []OrderQty
	Amount          int64 `json:"Amount" binding:"required"`
	Address 		Address
	Status 			string `json:"status" default:"pending"`
}

type Address struct {
	 gorm.Model
	 Addresses 		string `json:"addresses" binding:"required"`
	 OrderID 		int64 `json:"order_id"`
}

type OrderQty struct {
	gorm.Model
	CartID 		    int64 `json:"cart_id" binding:"required"`
	Quantity  		int64  `json:"quantity" binding:"required"`
	OrderID 		int64 `json:"order_id"`
}
func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Order{}, &Address{}, &OrderQty{})
}

func (c *Order) CreateOrder() *Order {
	db.NewRecord(c)
	db.Create(c)
	return c
}

func GetOrderById(Id int64) (*Order, *gorm.DB) {
	var getUser Order
	db := db.Where("ID=?", Id).Preload("OrderQuantity").Find(&getUser)
	return &getUser, db
}

