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
	dbo = config.GetDB()
	dbc.AutoMigrate(&Order{}, &Address{}, &OrderQty{})
}

func (c *Order) CreateOrder() *Order {
	dbc.NewRecord(c)
	dbc.Create(c)
	return c
}

func GetOrderById(Id int64) (*Order, *gorm.DB) {
	var getUser Order
	db := dbu.Where("ID=?", Id).Preload("OrderQuantity").Find(&getUser)
	return &getUser, db
}

